package route

import (
	"go_custom/controller"

	"github.com/gin-gonic/gin"
)

type Route struct{}

func (Route) Load(r *gin.Engine) {
	r.GET("/", controller.Index)
}
