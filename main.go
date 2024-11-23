package main

import (
	"github.com/gin-contrib/sessions"
	"go_custom/route"
	"go_custom/wrap/config"
	"go_custom/wrap/cookie"
	"go_custom/wrap/middleware"
	"go_custom/wrap/session"
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
	middleware.Load(r)

	// 设置cookie和session配置
	cookie.Load()
	store := session.Load()
	r.Use(sessions.Sessions("GlobalSession", store))

	// 构建路由
	route.Route.Load(route.Route{}, r)

	err = r.Run(":" + strconv.Itoa(config.Mapping.App.Port))
	if err != nil {

	}
}
