package main

import (
	"gin_work/route"
	"gin_work/wrap/config"
	"gin_work/wrap/cookie"
	"gin_work/wrap/middleware"
	"gin_work/wrap/session"
	"strconv"

	"github.com/gin-contrib/sessions"

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

	// 加载cookie和session配置
	cookie.Load()
	store := session.Load()
	r.Use(sessions.Sessions("GlobalSession", store))

	// 构建路由
	route.Route.Load(route.Route{}, r)

	err = r.Run(":" + strconv.Itoa(config.Mapping.App.Port))
	if err != nil {

	}
}
