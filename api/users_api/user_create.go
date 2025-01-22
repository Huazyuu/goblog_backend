package users_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/userServer"
)

type UserCreateRequest struct {
	Nickname string     `json:"nickname" binding:"required" msg:"请输入昵称"`
	UserName string     `json:"username" binding:"required" msg:"请输入用户名"`
	Password string     `json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3" msg:"请选择权限"`
}

// UserCreateView 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Param data body UserCreateRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UsersApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err := userServer.UserService{}.CreateUser(cr.UserName, cr.Nickname, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户创建成功 %s ", cr.UserName), c)
}
