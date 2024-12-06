package controller

import (
	"gin_work/common"
	"gin_work/extend/random"
	"gin_work/model"
	"gin_work/wrap/cache"
	"gin_work/wrap/response"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLoginForm struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required,alphanum,min=5,max=20"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required,alphanum,min=8,max=20"`
	Remember string `form:"remember" json:"remember" uri:"remember" xml:"remember" binding:"omitempty,lte=1"`
}

type CodeLoginForm struct {
	Phone    string `form:"phone" json:"phone" uri:"phone" xml:"phone" binding:"required,PhoneFormat"`
	Code     string `form:"code" json:"code" uri:"code" xml:"code" binding:"required,numeric,len=6"`
	Remember string `form:"remember" json:"remember" uri:"remember" xml:"remember" binding:"omitempty,lte=1"`
}

type SendCodeForm struct {
	Phone string `form:"phone" json:"phone" uri:"phone" xml:"phone" binding:"required,PhoneFormat"`
}

func Index(c *gin.Context) {
	str := "今天吃什么 "
	restaurants := []string{
		"稻花香快餐店",
		"每日一瞧猪脚饭",
		"价格昂贵鹅先生",
		"天天爆满木桶饭",
		"便宜少量烧鸭饭",
		"太二了螺狮粉",
		"13元自助餐",
		"黄焖说鸡不说吧",
	}
	num := rand.Intn(len(restaurants))
	str += "首选: " + restaurants[num]
	restaurants = append(restaurants[:num], restaurants[num+1:]...)
	num = rand.Intn(len(restaurants))
	str += "    次选: " + restaurants[num]
	c.JSON(http.StatusOK, gin.H{"result": str})
}

func UserLogin(c *gin.Context) {
	var login UserLoginForm
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(response.RequestFail(err.Error()))
		return
	}

	var user *model.User
	user.FindByUsername(login.Username)
	if user == nil {
		c.JSON(response.RequestFail("用户账号或密码错误"))
		return
	}
	if err := common.CheckPwd(login.Password, user.Password, user.Salt); err != nil {
		c.JSON(response.RequestFail(err.Error()))
		return
	}

	// 登录更新
	UserLoginUpdate(user, c.ClientIP())

	// 生成token
	if err := createToken(c, user); err != nil {
		c.JSON(response.Fail(err.Error()))
		return
	}

	c.JSON(response.Success())
	return
}

func CodeLogin(c *gin.Context) {
	var login CodeLoginForm
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(response.RequestFail(err.Error()))
		return
	}

	codeTime, err1 := cache.Get("code_time_" + login.Phone)
	code, err2 := cache.Get("code_" + login.Phone)
	if err1 != nil {
		c.JSON(response.Fail(err1.Error()))
		return
	} else if err2 != nil {
		c.JSON(response.Fail(err2.Error()))
		return
	} else if codeTime == nil {
		c.JSON(response.RequestFail("验证码已过期,请重新验证"))
		return
	} else if code != login.Code {
		c.JSON(response.RequestFail("验证码错误,请重新验证"))
		return
	}

	var user *model.User
	user.FindByPhone(login.Phone)

	// 登录更新
	UserLoginUpdate(user, c.ClientIP())

	// 生成token
	if err := createToken(c, user); err != nil {
		c.JSON(response.Fail(err.Error()))
		return
	}

	c.JSON(response.Success())
	return
}

func SendCode(c *gin.Context) {
	var codeForm *SendCodeForm
	if err := c.ShouldBind(&codeForm); err != nil {
		c.JSON(response.RequestFail(err.Error()))
	}

	codeTime, _ := cache.Get("code_time_" + codeForm.Phone)
	if codeTime != nil {
		remainTime := codeTime.(time.Time).Unix() - time.Now().Unix()
		c.JSON(response.RequestFail("验证码 " + strconv.Itoa(int(remainTime)) + " 秒后才能再次发送"))
		return
	}

	randNumber := random.RandNum(6)
	err1 := cache.Set("code_"+codeForm.Phone, randNumber, 60)
	err2 := cache.Set("code_time_"+codeForm.Phone, time.Now(), 60)
	if err1 != nil || err2 != nil {
		c.JSON(response.Fail("验证码生成失败,请重试"))
		return
	}
	// 发送验证码

	c.JSON(response.Success())
	return
}

func createToken(c *gin.Context, user *model.User) error {
	accessToken, err := common.RefreshToken(user.Uuid, c.Request.Host, c.ClientIP())
	if err != nil {
		return err
	}
	c.Header("Authorization", accessToken)
	_ = cache.Set(user.Uuid, accessToken)
	return nil
}
