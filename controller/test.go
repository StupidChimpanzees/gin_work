package controller

import (
	"gin_work/model"
	"gin_work/wrap/cache"
	"gin_work/wrap/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TestParams struct {
	Id      int    `form:"id" json:"id" uri:"id" xml:"id"`
	OrderNo string `form:"orderNo" json:"orderNo" uri:"orderNo" xml:"orderNo"`
	Count   int    `form:"count" json:"count" uri:"count" xml:"count"`
}

func Test(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	i, _ := strconv.Atoi(id)
	storagePlanModel := model.StoragePlanModel{}
	planInfo, err := storagePlanModel.FindJoinById(i)
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

	get, err := cache.Get("test")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := map[string]any{"cache": get, "bind": params}
	c.JSON(response.Success(result))
}
