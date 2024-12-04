package middleware

import (
	"fmt"
	"gin_work/middleware"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	var GlobalMiddle *middleware.GlobalMiddleware
	globalValue := reflect.ValueOf(GlobalMiddle)
	fmt.Println(globalValue.NumMethod())
	for i := 0; i < globalValue.NumMethod(); i++ {
		funcResult := globalValue.Method(i).Call(nil)
		r.Use(funcResult[0].Interface().(gin.HandlerFunc))
	}
}
