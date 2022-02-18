package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"ihomeMs/controller"
	"ihomeMs/utils"
	"net/http"
	"time"
)

// Filter 路由过滤器
func Filter(ctx *gin.Context) {
	//登录校验
	session := sessions.Default(ctx)
	userName := session.Get("userName")
	resp := make(map[string]interface{})
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		ctx.JSON(http.StatusOK, resp)
		ctx.Abort()
		return
	}

	//计算这个业务耗时
	fmt.Println("next之前打印", time.Now())

	//执行函数
	ctx.Next()

	fmt.Println("next之后打印....")
}

func main() {
	//初始化路由
	router := gin.Default()

	//初始化redis容器,存储session数据
	store, _ := redis.NewStore(10, "tcp", "47.94.195.58:6379", "123456", []byte("session"))
	router.Use(sessions.Sessions("mySession", store))

	//静态路由
	router.Static("/home", "./view")

	r1 := router.Group("/api/v1.0")
	{
		//路由规范
		r1.GET("/areas", controller.GetArea)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:mobile", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
		r1.POST("/sessions", controller.PostLogin)
		r1.GET("/session", controller.GetSession)

		//路由过滤器   登录的情况下才能执行一下路由请求
		r1.Use(Filter)
		r1.DELETE("/session", controller.DeleteSession)
		r1.GET("/user", controller.GetUserInfo)
		r1.PUT("/user/name", controller.PutUserInfo)
		r1.POST("/user/avatar", controller.PostAvatar)
		r1.POST("/user/auth", controller.PutUserAuth)
		r1.GET("/user/auth", controller.GetUserInfo)
		//获取已发布房源信息
		r1.GET("/user/houses", controller.GetUserHouses)
		//发布房源
		r1.POST("/houses", controller.PostHouses)
		//添加房源图片
		r1.POST("/houses/:id/images", controller.PostHousesImage)
		//展示房屋详情
		r1.GET("/houses/:id", controller.GetHouseInfo)
		//展示首页轮播图
		r1.GET("/house/index", controller.GetIndex)
		//搜索房屋
		r1.GET("/houses", controller.GetHouses)
		//下订单
		r1.POST("/orders", controller.PostOrders)
		//获取订单
		r1.GET("/user/orders", controller.GetUserOrder)
		//同意/拒绝订单
		r1.PUT("/orders/:id/status", controller.PutOrders)
		//发表评价
		r1.PUT("/orders/:id/comment", controller.PutComment)
	}

	router.Run("0.0.0.0:8090")
}
