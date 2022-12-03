package router

import (
	"draw-service/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 用户登录
	r.POST("/login", service.Login)
	// 发送验证验证码
	r.POST("/sendSMS", service.SendSMS)
	r.POST("/register", service.Register)

	return r
}
