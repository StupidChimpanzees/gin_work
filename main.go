package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"workspace/src/gin_custom/extend/config"
)

func main() {
	r := gin.Default()

	// 加载配置
	err := config.Info.Load("config.yaml")
	if err != nil {
		return
	}
	fmt.Println("----------------")
	fmt.Println(config.Instance.Set("app.version", "1.1.0"))

	// 加载中间件

	// 构建路由

	err = r.Run("")
	if err != nil {

	}
}
