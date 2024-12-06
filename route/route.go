package route

import (
	"gin_work/controller"

	"github.com/gin-gonic/gin"
)

type Route struct{}

func (*Route) Home(r *gin.Engine) {
	r.GET("/index", controller.Index)
}

func (*Route) Login(r *gin.Engine) {
	r.POST("/login", controller.UserLogin)
	r.GET("/send_code", controller.SendCode)
	r.POST("/code_login", controller.CodeLogin)
}
