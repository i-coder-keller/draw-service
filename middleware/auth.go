package middleware

import (
	"draw-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "用户未登录",
			})
			c.Abort()
			return
		}
		token = strings.Split(token, "Bearer ")[1]
		jwtData, err := utils.AnalyseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "用户未登录",
			})
			c.Abort()
			return
		}
		if time.Now().UnixMilli() > jwtData.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "登录过期",
			})
			c.Abort()
			return
		}
		c.Set("auth", jwtData)
		c.Next()
	}
}
