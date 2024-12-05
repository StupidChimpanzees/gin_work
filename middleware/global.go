package middleware

import (
	"gin_work/common"
	"gin_work/extend/jwt"
	"gin_work/wrap/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GlobalMiddleware struct{}

func (*GlobalMiddleware) Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	}
}

// Cors 开启跨域请求
func (*GlobalMiddleware) Cors() gin.HandlerFunc {
	return cors.Default()
}

func (*GlobalMiddleware) Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !common.ExceptUrl(c, "/login", "/sign_in") {
			accessToken := c.GetHeader("access_token")
			refreshToken := c.GetHeader("refresh_token")
			var claims *jwt.TokenClaims
			claims, accessB, refreshB := common.CheckToken(accessToken, refreshToken, c.Request.Host, c.ClientIP())
			if !accessB && !refreshB {
				c.AbortWithStatusJSON(response.Fail(http.StatusBadRequest, "用户认证信息已过期"))
			} else if !accessB && refreshB {
				accessToken, refreshToken, err := common.RefreshNewToken(claims.ID, claims.Subject, claims.Ip)
				if err != nil {
					c.AbortWithStatusJSON(response.Fail(http.StatusBadRequest, "用户认证信息无法更新"))
				}
				c.Header("access_token", accessToken)
				c.Header("refresh_token", refreshToken)
			}
			c.Set("token", claims)
		}
		c.Next()
	}
}
