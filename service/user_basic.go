package service

import (
	"draw-service/models"
	"draw-service/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

// 发送验证码
func SendSMS(c *gin.Context) {
	type Email struct {
		Email string `json:"email"`
	}
	request := &Email{}
	c.ShouldBindJSON(request)
	if request.Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	count, err := models.GetUserBasicByEmail(request.Email)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前邮箱已被注册",
		})
		return
	}
	err = utils.SendCode(request.Email, "1234")
	if err != nil {
		log.Printf("[ERROR]:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}
