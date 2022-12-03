package service

import (
	"draw-service/models"
	"draw-service/utils"
	"log"
	"net/http"
	"time"

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
	token, err := utils.GenerateToken(ub.Identity, ub.Email, ub.Account)
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
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "服务器错误",
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前邮箱已被注册",
		})
		return
	}
	expires := time.Now().UnixMilli()
	ev, err := models.GetEmailValidationByEmail(request.Email)
	log.Println(expires, ev.SendTime)
	if expires-ev.SendTime < 60000 {
		log.Printf("[短时间重复发送]")
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码发送频繁",
		})
		return
	}
	if err != nil {
		log.Println("首次生成", err)
		code := utils.RandomCode()
		err = models.InsertCodeAndEmailAndExpires(code, request.Email, expires+300000, expires)
		if err != nil {
			log.Printf("[插入新数据失败]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "验证码发送失败",
			})
			return
		}
		err = utils.SendCode(request.Email, code)
		if err != nil {
			log.Printf("[发送验证码失败]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "验证码发送失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证码发送成功",
		})
		return
	}
	if expires >= ev.Expires {
		// 时间过期重新生成
		log.Println("时间过期重新生成")
		code := utils.RandomCode()
		err = models.UpdateCodeAndExpiresByEmail(code, expires, ev.Email)
		if err != nil {
			log.Printf("[ERROR]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "验证码发送失败",
			})
			return
		}
		err = utils.SendCode(request.Email, code)
		if err != nil {
			log.Printf("[ERROR]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "验证码发送失败",
			})
			return
		}
	}
	if expires < ev.Expires {
		// 没有过期
		log.Println("时间未过期直接发送")
		err = utils.SendCode(request.Email, ev.Code)
		if err != nil {
			log.Printf("[ERROR]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "验证码发送失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证码发送成功",
		})
	}
}

// 注册
func Register(c *gin.Context) {
	type register struct {
		Account  string `json:"account"`
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Code     string `json:"code"`
		Password string `json:"password"`
	}
	request := new(register)
	c.ShouldBindJSON(request)
	if request.Account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号不能为空",
		})
		return
	}
	if request.Nickname == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号不能为空",
		})
		return
	}
	if request.Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	if request.Code == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不能为空",
		})
		return
	}
	if request.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "密码不能为空",
		})
		return
	}
	avatar := "http://drawcdn.liuyongzhi.cn/default-avatar.png"
	ev, err := models.GetEmailValidationByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}
	if ev.Code != request.Code {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}
	ub, err := models.GetUserBasicByAccountAndEmail(request.Account, request.Email)
	if err != nil {
		// 未注册
		err = models.InsertUserBasic(request.Account, request.Email, request.Nickname, utils.GetMd5(request.Password), avatar, time.Now().UnixMilli(), time.Now().UnixMilli())
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "注册失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "注册成功",
		})
		return
	}
	if ub.Account != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户已被注册",
		})
	}
}
