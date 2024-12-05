package common

import "github.com/gin-gonic/gin"

func ExceptUrl(c *gin.Context, args ...string) bool {
	path := c.Request.URL.Path
	for k := range args {
		if args[k] == path {
			return true
		}
	}
	return false
}
