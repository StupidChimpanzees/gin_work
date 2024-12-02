package controller

import (
	"gin_work/model"
	"gin_work/wrap/cache"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TestParams struct {
	Id      int    `form:"id" json:"id" uri:"id" xml:"id"`
	OrderNo string `form:"orderNo" json:"orderNo" uri:"orderNo" xml:"orderNo"`
	Count   int    `form:"count" json:"count" uri:"count" xml:"count"`
}

func Test(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	i, _ := strconv.Atoi(id)
	planInfo, err := model.StoragePlanModel.FindJoinById(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = cache.Set("test", planInfo, 5)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var params TestParams
	err = cache.BindGet("test", &params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	get, err := cache.AnyGet("test")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := map[string]any{"plan_info": planInfo, "cache": get, "bind": params}
	c.JSON(http.StatusOK, gin.H{"status": "200", "data": result})
}
