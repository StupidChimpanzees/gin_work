package route

import (
	"gin_work/route"
	"github.com/gin-gonic/gin"
	"reflect"
)

func Load(r *gin.Engine) {
	routeType := reflect.ValueOf(route.Route{})
	routeValue := []reflect.Value{reflect.ValueOf(r)}
	for i := 0; i < routeType.NumMethod(); i++ {
		routeType.Method(i).Call(routeValue)
	}
}
