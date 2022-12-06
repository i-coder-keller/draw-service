package router

import (
	"draw-service/middleware"
	"draw-service/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 用户登录
	r.POST("/login", service.Login)
	// 发送验证验证码
	r.POST("/sendSMS", service.SendSMS)
	// 用户注册
	r.POST("/register", service.Register)

	AuthGrounp := r.Group("/api", middleware.AuthMiddleware())
	AuthGrounp.POST("/auth", service.Auth)
	return r
}
