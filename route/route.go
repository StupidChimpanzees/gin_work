package route

import (
	"github.com/gin-gonic/gin"
	"go_custom/controller"
)

type Route struct{}

func (Route) Load(r *gin.Engine) {
	r.GET("/", controller.Index)
}
