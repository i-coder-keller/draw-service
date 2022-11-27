package service

import (
	"draw-service/models"
	"draw-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	type postJson struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	request := &postJson{}
	c.ShouldBindJSON(request)
	if request.Account == "" || request.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": gin.H{},
			"msg":  "用户名或密码不能为空",
		})
		return
	}
	ub, err := models.GetUserBasicByAccountAndPassword(request.Account, utils.GetMd5(request.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": gin.H{},
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := utils.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"token":    token,
			"userInfo": ub,
		},
		"msg": "登录成功",
	})
}
