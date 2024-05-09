package userServer

import (
	"gvb_server/service/redisServer"
	"gvb_server/utils/jwt"
	"time"
)

func (UserService) Logout(claims *jwt.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redisServer.Logout(token, diff)
}
