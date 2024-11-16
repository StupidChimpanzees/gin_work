package main

import (
	"github.com/gin-gonic/gin"
	"go_custom/extend/config"
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

	err = r.Run(config.Get("app.port").(string))
	if err != nil {

	}
}
