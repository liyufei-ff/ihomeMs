package handler

import (
	"context"
	"strconv"

	"fmt"
	"ihomeMs/service/userOrder/model"
	userOrder "ihomeMs/service/userOrder/proto/userOrder"
	"ihomeMs/service/userOrder/utils"
)

type UserOrder struct{}

// CreateOrder 创建订单
func (e *UserOrder) CreateOrder(ctx context.Context, req *userOrder.Request, rsp *userOrder.Response) error {
	//判断房间是否被抢订
	is := model.IsRushOrder(req.StartDate, req.EndDate, req.HouseId)

	if is {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	} else {
		//获取到相关数据,插入到数据库
		orderId, err := model.InsertOrder(req.HouseId, req.StartDate, req.EndDate, req.UserName)
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
			return nil
		}
		rsp.Errno = utils.RECODE_OK
		rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		var orderData userOrder.OrderData
		orderData.OrderId = strconv.Itoa(orderId)

		rsp.Data = &orderData
	}

	return nil
}

func (e *UserOrder) GetOrderInfo(ctx context.Context, req *userOrder.GetReq, resp *userOrder.GetResp) error {
	//要根据传入数据获取订单信息   mysql
	respData, err := model.GetOrderInfo(req.UserName, req.Role)
	if err != nil {
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	var getData userOrder.GetData
	getData.Orders = respData
	resp.Data = &getData

	return nil
}

func (e *UserOrder) UpdateStatus(ctx context.Context, req *userOrder.UpdateReq, resp *userOrder.UpdateResp) error {
	//根据传入数据,更新订单状态
	err := model.UpdateStatus(req.Action, req.Id, req.Reason)
	if err != nil {
		fmt.Println("更新订单装填错误", err)
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	return nil
}

func (e *UserOrder) UpdateOrderComment(ctx context.Context, req *userOrder.OrderCommentReq, resp *userOrder.OrderCommentResp) error {
	//更新订单的评价信息
	var err error
	err = model.UpdateOrderComment(req.OrderId, req.Comment)
	//完成订单信息评价之后 房屋的入住次数加一
	err = model.AddOrderCount(req.OrderId)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = err.Error()
	} else {
		resp.Errno = utils.RECODE_OK
		resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	}
	return nil
}
