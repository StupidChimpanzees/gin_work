package controller

import (
	"gin_work/wrap/response"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type UserLogin struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required,alphanum,min=5,max=20"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required,alphanum,min=8,max=20"`
	Code     string `form:"code" json:"code" uri:"code" xml:"code" binding:"omitempty,required,alphanum,len=6"`
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

func Login(c *gin.Context) {
	var user UserLogin
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(response.Fail(http.StatusBadRequest, err.Error()))
	}

}
