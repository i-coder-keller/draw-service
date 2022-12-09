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

	AuthGroup := r.Group("/api", middleware.AuthMiddleware())
	AuthGroup.POST("/auth", service.Auth)
	AuthGroup.POST("/projects", service.Projects)
	AuthGroup.POST("/createProject", service.CreateProject)
	AuthGroup.POST("/invite", service.InviteOfProject)
	return r
}
