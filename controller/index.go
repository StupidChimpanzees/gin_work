package controller

import (
	"gin_work/common"
	"gin_work/model"
	"gin_work/wrap/response"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required,alphanum,min=5,max=20"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required,alphanum,min=8,max=20"`
	Code     string `form:"code" json:"code" uri:"code" xml:"code" binding:"omitempty,alphanum,len=6"`
	Remember string `form:"remember" json:"remember" uri:"remember" xml:"remember" binding:"omitempty,lte=1"`
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
	var login LoginForm
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(response.Fail(http.StatusBadRequest, err.Error()))
		return
	}

	var user *model.User
	user.FindByUsername(login.Username)
	if user == nil {
		c.JSON(response.Fail(http.StatusBadRequest, "用户账号或密码错误"))
		return
	}
	if err = common.CheckPwd(login.Password, user.Password, user.Salt); err != nil {
		c.JSON(response.Fail(http.StatusBadRequest, err.Error()))
		return
	}

	// 登录更新
	UserLoginUpdate(user, c.ClientIP())

	// 生成token
	accessToken, err := common.RefreshToken(user.Uuid, c.Request.Host, c.ClientIP())
	if err != nil {
		c.JSON(response.Fail(http.StatusBadRequest, err.Error()))
	}
	c.Header("access_token", accessToken)

}
