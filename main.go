package main

import (
	"github.com/gin-gonic/gin"
	"workspace/src/framework/extend/config"
)

func main() {
	r := gin.Default()

	// 加载配置
	err := config.Info.Load("config.yaml")
	if err != nil {
		return
	}

	// 加载中间件

	// 构建路由

	err = r.Run("")
	if err != nil {

	}
}
