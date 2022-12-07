package service

import (
	"draw-service/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Projects(c *gin.Context) {
	Identity := c.GetString("identity")
	result, newErr := models.FindAllProjectByOwnerIdentity(Identity)

	log.Println(newErr)
	if newErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "用户异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  "查询成功",
	})
}
