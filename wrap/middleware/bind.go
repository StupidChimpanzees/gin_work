package middleware

import "github.com/gin-gonic/gin"

func init() {

}

func Bind(r *gin.Engine, global func() gin.HandlerFunc) {

	r.Use(global())
}
