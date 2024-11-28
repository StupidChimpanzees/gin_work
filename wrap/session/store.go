package session

import (
	"gin_work/wrap/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var c *gin.Context

func Load() cookie.Store {
	return cookie.NewStore([]byte(getConfig().Secret))
}

func getConfig() config.SessionConfiguration {
	return config.Mapping.Session
}

func SetContext(context *gin.Context) {
	c = context
}

func Set(name string, value interface{}) error {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: getConfig().Expire})
	session.Set(name, value)
	return session.Save()
}

func Get(name string) interface{} {
	session := sessions.Default(c)
	return session.Get(name)
}

func Delete(name string) {
	session := sessions.Default(c)
	session.Delete(name)
}

func Clear() {
	session := sessions.Default(c)
	session.Clear()
}
