package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/redisServer"
	"gvb_server/utils/jwt"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 是否在redis中
		if redisServer.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}
		if redisServer.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
