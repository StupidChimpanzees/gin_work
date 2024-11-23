package middleware

import (
	"github.com/gin-gonic/gin"
)

type GlobalMiddleware struct {
}

var GM GlobalMiddleware

func (GlobalMiddleware) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (GlobalMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
