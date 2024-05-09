package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils/jwt"
)

func (UsersApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	token := c.Request.Header.Get("token")

	err := service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}

	res.OkWithMessage("注销成功", c)

}
