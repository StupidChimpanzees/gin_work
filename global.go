package main

import "github.com/gin-gonic/gin"

type GlobalMiddleware struct {
	group map[string]func() gin.HandlerFunc
}

func Login() {}

func Auth() {}
