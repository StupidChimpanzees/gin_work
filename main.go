package main

import (
	"go_custom/extend/config"
	"strconv"

	"github.com/gin-gonic/gin"

type User struct {
	Name string
	Age  int
}

func main() {
	r := gin.Default()

	// 加载配置
	err := config.Load("config.yaml")
	if err != nil {
		return
	}

	// 构建路由
	route.Route.Load(route.Route{}, r)

	// 构建路由
	route.Route.Load(route.Route{}, r)

	err = r.Run(":" + strconv.Itoa(config.Mapping.App.Port))
	if err != nil {

	}
}
