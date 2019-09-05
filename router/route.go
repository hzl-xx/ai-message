package router

import (
	"github.com/gin-gonic/gin"
	"messageserver/controller"
	"messageserver/controller/v1"
	"messageserver/middleware"
)

var (
	Router *gin.Engine
)

func init() {
	Router = gin.Default()
	//gin.SetMode("debug")
	initRoute()
}

func initRoute() {

	wechat := &v1.WechatController{}
	token := &controller.TokenController{}
	uc := &controller.UserController{}
	Router.GET("/key", token.GetKey)
	Router.GET("/auth", token.GetToken)
	Router.POST("/send/message", wechat.SendMessage)
	v1 := Router.Group("v1").Use(middleware.JWT())
	{
		v1.POST("/send/message", wechat.SendMessage)
		v1.GET("/send/message1", uc.GetAuth)

	}
}