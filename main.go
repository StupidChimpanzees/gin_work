package main

import (
	"go_custom/route"
	"go_custom/wrap/config"
	"go_custom/wrap/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载配置
	err := config.Load("config.yaml")
	if err != nil {
		return
	}

	// 加载全局中间件
	middleware.Bind(r)

	// 构建路由
	route.Route.Load(route.Route{}, r)

	err = r.Run(":" + strconv.Itoa(config.Mapping.App.Port))
	if err != nil {

	}
}
