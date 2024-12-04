package route

import (
	"gin_work/controller"

	"github.com/gin-gonic/gin"
)

type Route struct{}

func (*Route) Load(r *gin.Engine) {
	r.GET("/index", controller.Index)
	r.GET("/test", controller.Test)
}

func (*Route) Login(r *gin.Engine) {
	r.POST("/login", controller.Login)
}
