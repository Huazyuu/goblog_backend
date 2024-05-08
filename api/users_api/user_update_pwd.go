package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
)

type UserUpdatePwdReq struct {
	NewPwd string `json:"new_pwd"`
	OldPwd string `json:"old_pwd"`
}

func (UsersApi) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var cr UserUpdatePwdReq
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	flag, _ := pwd.CheckPwd(user.Password, cr.OldPwd)
	if !flag {
		res.FailWithMessage("原密码输入错误", c)
		return
	}

	hashPwd, _ := pwd.HashPwd(cr.NewPwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("修改密码失败", c)
		return
	}
	res.OkWithMessage("修改密码成功", c)

}
