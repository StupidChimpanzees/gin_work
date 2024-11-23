package cookie

import (
	"github.com/gin-gonic/gin"
	"go_custom/wrap/config"
)

type cookiePack struct {
	conf config.CookieConfiguration
}

var cPack *cookiePack

func Load() {
	cPack.conf = getConfig()
}

func getConfig() config.CookieConfiguration {
	return config.Mapping.Cookie
}

func Set(c *gin.Context, name, value string) {
	cc := cPack.conf
	c.SetCookie(name, value, cc.Expire, cc.Path, cc.Domain, cc.Secure, cc.HttpOnly)
}

func Get(c *gin.Context, name string) (string, error) {
	return c.Cookie(name)
}

func Delete(c *gin.Context, name string) {
	cc := cPack.conf
	c.SetCookie(name, "", -1, cc.Path, cc.Domain, cc.Secure, cc.HttpOnly)
}
