package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/desensitize"
	"gvb_server/utils/jwt"
)

func (UsersApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 分页查询
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, cnt, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员可见username
			user.UserName = ""
		}
		// 手机号脱敏
		user.Tel = desensitize.DesensitizationTel(user.Tel)
		user.Email = desensitize.DesensitizationEmail(user.Email)
		users = append(users, user)
	}
	res.OkWithList(users, cnt, c)
}
