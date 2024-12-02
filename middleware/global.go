package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GlobalMiddleware struct{}

func (*GlobalMiddleware) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (*GlobalMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Cors 开启跨域请求
func (*GlobalMiddleware) Cors() gin.HandlerFunc {
	return cors.Default()
}
