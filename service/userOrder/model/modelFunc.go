package model

import (
	"fmt"
	userOrder "ihomeMs/service/userOrder/proto/userOrder"
	"strconv"
	"time"
)

type UserData struct {
	Id int
}

// IsRushOrder 判断房间是否被抢订以及租房天数必须>=minDays <=maxDays
func IsRushOrder(startDate, endDate, houseId string) bool {
	//将时间字符串转换为时间类型
	hid, _ := strconv.Atoi(houseId)
	sDate, _ := time.Parse("2006-01-02", startDate)
	eDate, _ := time.Parse("2006-01-02", endDate)

	//租房天数
	sub := eDate.Sub(sDate)
	days := int(sub.Hours() / 24)
	houseInfo := GetHouseInfoByHid(hid)
	if days < houseInfo.Min_days || days > houseInfo.Max_days { //租房天数超范围
		fmt.Println("租房天数超范围", days, houseInfo.Min_days, houseInfo.Max_days)
		return true
	}

	var startCount int
	var endCount int

	db := GlobalDB.Model(new(OrderHouse)).Where("house_id = ?", hid).
		Where("status = ?", "WAIT_ACCEPT")

	db.Where("begin_date BETWEEN ? AND ?", sDate, eDate).Count(&startCount)
	db.Where("end_date BETWEEN ? AND ?", sDate, eDate).Count(&endCount)
	if startCount == 0 && endCount == 0 {
		return false
	} else {
		return true
	}
}

// GetHouseInfoByHid 通过房屋ID获取到房屋信息
func GetHouseInfoByHid(hid int) House {
	var houseInfo House
	GlobalDB.Where("id = ?", hid).Find(&houseInfo)
	return houseInfo
}

func InsertOrder(houseId, beginDate, endDate, userName string) (int, error) {
	//获取插入对象
	var order OrderHouse

	//给对象赋值
	hid, _ := strconv.Atoi(houseId)
	order.HouseId = uint(hid)

	//把string类型的时间转换为time类型
	bDate, _ := time.Parse("2006-01-02", beginDate)
	order.Begin_date = bDate

	eDate, _ := time.Parse("2006-01-02", endDate)
	order.End_date = eDate

	var userData UserData
	if err := GlobalDB.Raw("select id from user where name = ?", userName).Scan(&userData).Error; err != nil {
		fmt.Println("获取用户数据错误", err)
		return 0, err
	}

	//获取days
	dur := eDate.Sub(bDate)
	order.Days = int(dur.Hours()) / 24
	order.Status = "WAIT_ACCEPT"

	//房屋的单价和总价
	var house House
	GlobalDB.Where("id = ?", hid).Find(&house).Select("price")
	order.House_price = house.Price
	order.Amount = house.Price * order.Days

	order.UserId = uint(userData.Id)
	if err := GlobalDB.Create(&order).Error; err != nil {
		fmt.Println("插入订单失败", err)
		return 0, err
	}
	return int(order.ID), nil
}

// GetOrderInfo 获取房东订单如何实现?
func GetOrderInfo(userName, role string) ([]*userOrder.OrdersData, error) {
	//最终需要的数据
	var orderResp []*userOrder.OrdersData
	//获取当前用户的所有订单
	var orders []OrderHouse

	var userData UserData
	//用原生查询的时候,查询的字段必须跟数据库中的字段保持一直
	GlobalDB.Raw("select id from user where name = ?", userName).Scan(&userData)

	//查询租户的所有的订单
	if role == "custom" {
		if err := GlobalDB.Where("user_id = ?", userData.Id).Find(&orders).Error; err != nil {
			fmt.Println("获取当前用户所有订单失败")
			return nil, err
		}
	} else {
		//查询房东的订单  以房东视角来查看订单
		var houses []House
		GlobalDB.Where("user_id = ?", userData.Id).Find(&houses)

		for _, v := range houses {
			var tempOrders []OrderHouse
			GlobalDB.Model(&v).Related(&tempOrders)

			orders = append(orders, tempOrders...)
		}
	}

	//循环遍历一下orders
	for _, v := range orders {
		var orderTemp userOrder.OrdersData
		orderTemp.OrderId = int32(v.ID)
		orderTemp.EndDate = v.End_date.Format("2006-01-02")
		orderTemp.StartDate = v.Begin_date.Format("2006-01-02")
		orderTemp.Ctime = v.CreatedAt.Format("2006-01-02")
		orderTemp.Amount = int32(v.Amount)
		orderTemp.Comment = v.Comment
		orderTemp.Days = int32(v.Days)
		orderTemp.Status = v.Status

		//关联house表
		var house House
		GlobalDB.Model(&v).Related(&house).Select("index_image_url", "title")
		orderTemp.ImgUrl = "http://47.94.195.58:8888/" + house.Index_image_url
		orderTemp.Title = house.Title

		orderResp = append(orderResp, &orderTemp)
	}
	return orderResp, nil
}

// UpdateStatus 更新订单状态
func UpdateStatus(action, id, reason string) error {
	db := GlobalDB.Model(new(OrderHouse)).Where("id = ?", id)

	if action == "accept" {
		//标示房东同意订单
		return db.Update("status", "WAIT_COMMENT").Error
	} else {
		//表示房东不同意订单  如果拒单把拒绝的原因写到comment中
		return db.Updates(map[string]interface{}{"status": "REJECTED", "comment": reason}).Error
	}
}

// UpdateOrderComment 更新订单评价信息
func UpdateOrderComment(oid, comment string) error {
	orderId, _ := strconv.Atoi(oid)
	return GlobalDB.Model(new(OrderHouse)).
		Where("id = ?", uint(orderId)).Updates(map[string]interface{}{"status": "COMPLETE", "comment": comment}).Error
}

// AddOrderCount 房屋入住次数加一
func AddOrderCount(oid string) error {
	orderId, _ := strconv.Atoi(oid)
	var err error
	var orderInfo OrderHouse
	var houseInfo House
	err = GlobalDB.Where("id = ?", uint(orderId)).Select("house_id").Find(&orderInfo).Error
	err = GlobalDB.Where("id = ?", orderInfo.HouseId).Select("order_count").Find(&houseInfo).Error
	err = GlobalDB.Model(new(House)).Where("id = ?", orderInfo.HouseId).Update("order_count", houseInfo.Order_count+1).Error

	return err
}
