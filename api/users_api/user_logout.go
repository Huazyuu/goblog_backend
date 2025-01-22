package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils/jwt"
)

// LogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Param token header string  true  "token"
// @Router /api/logout [post]
// @Produce json
// @Success 200 {object} res.Response{}
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
