package main

import (
	"github.com/gin-gonic/gin"
	"go_custom/extend/config"
	"go_custom/route"
	"strconv"
)

func main() {
	r := gin.Default()

	// 加载配置
	err := config.Load("config.yaml")
	if err != nil {
		return
	}

	// 加载中间件

	// 构建路由
	route.Route.Load(route.Route{}, r)

	err = r.Run(strconv.Itoa(config.Mapping.App.Port))
	if err != nil {

	}
}
