package controller

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func Index(c *gin.Context) {
	str := "今天吃什么 "
	restaurants := []string{
		"周杰伦稻花香快餐店",
		"每日一瞧猪脚饭",
		"价格昂贵鹅先生",
		"天天爆满木桶饭",
		"便宜少量烧鸭饭",
		"太二了面条店",
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
