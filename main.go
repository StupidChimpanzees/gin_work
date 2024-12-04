package main

import (
	"gin_work/wrap/config"
	"gin_work/wrap/cookie"
	"gin_work/wrap/middleware"
	"gin_work/wrap/route"
	"gin_work/wrap/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	r := handler()

	// autotls.Run(r, "example1.com")
	err := r.Run(":" + strconv.Itoa(config.Mapping.App.Port))
	if err != nil {
		log.Fatalf("error info: " + err.Error())
	}
}

func handler() *gin.Engine {
	r := gin.Default()

	// 加载配置
	err := config.Load("config.yaml")
	if err != nil {
		panic("Config file load error")
	}

	// 加载全局中间件
	middleware.Load(r)

	// 加载cookie和session配置
	cookie.Load()
	store := session.Load()
	r.Use(sessions.Sessions("GlobalSession", store))

	// 加载view配置
	// 目录下必须有.html文件才能使用
	// view.Load(r)

	// 构建路由
	route.Load(r)

	_ = r.SetTrustedProxies(nil)

	return r
}
