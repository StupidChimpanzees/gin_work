package controller

import (
	"gin_work/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TestParams struct {
	Id int `form:"id" json:"id" uri:"id" xml:"id" binding:"required"`
}

func Test(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	i, _ := strconv.Atoi(id)
	planInfo, err := model.StoragePlanModel.FindJoinById(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "data": planInfo})
}
