package middleware

import (
	"gin_work/common"
	"gin_work/extend/jwt"
	"gin_work/wrap/cache"
	"gin_work/wrap/response"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
		if !common.ExceptUrl(c, "/sign_in", "/login", "/send_code", "/code_login") {
			accessToken := c.GetHeader("Authorization")
			var claims *jwt.TokenClaims
			claims, err := common.CheckToken(accessToken, c.Request.Host, c.ClientIP())
			if err == nil {
				// 用户唯一token是否有效
				localToken, _ := cache.Get(claims.ID)
				if localToken.(string) != accessToken {
					c.AbortWithStatusJSON(response.Fail(http.StatusBadRequest, "用户认证信息已失效"))
					return
				}
			} else if err.Error() == jwt.ExpiresErr {
				// 未超出最大刷新时间
				accessToken, err := common.RefreshToken(claims.ID, claims.Subject, claims.Ip)
				if err != nil {
					c.AbortWithStatusJSON(response.Fail(http.StatusBadRequest, "用户认证信息无法更新"))
					return
				}
				c.Header("Authorization", accessToken)
				_ = cache.Set(claims.ID, accessToken)
			} else if claims == nil {
				c.AbortWithStatusJSON(response.Fail(http.StatusBadRequest, "用户认证信息已过期"))
				return
			}
			c.Set("token", claims)
			// c.MustGet("token")
		}
		c.Next()
	}
}
