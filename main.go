package main

import (
	"fmt"
	"go_custom/extend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载配置

	err := config.Load("config.yaml")
	if err != nil {
		return
	}
	fmt.Println("----------------")
	fmt.Println(config.Set("app.version", "1.1.0"))
	fmt.Println(config.Get("database.port"))

	// 加载中间件

	// 构建路由

	err = r.Run("")
	if err != nil {

	}
}
