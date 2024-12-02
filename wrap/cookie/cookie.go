package cookie

import (
	"gin_work/wrap/config"

	"github.com/gin-gonic/gin"
)

var cPack config.CookieConfiguration

func Load() {
	cPack = getConfig()
}

func getConfig() config.CookieConfiguration {
	return config.Mapping.Cookie
}

func Set(c *gin.Context, name, value string) {
	cc := cPack
	c.SetCookie(name, value, cc.Expire, cc.Path, cc.Domain, cc.Secure, cc.HttpOnly)
}

func Get(c *gin.Context, name string) (string, error) {
	return c.Cookie(name)
}

func Delete(c *gin.Context, name string) {
	cc := cPack
	c.SetCookie(name, "", -1, cc.Path, cc.Domain, cc.Secure, cc.HttpOnly)
}
