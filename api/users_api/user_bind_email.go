package users_api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"非法邮箱"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

var emailFirstReq string

func (UsersApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 绑定邮箱 第一次发送code
	var cr BindEmailRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)

	if cr.Code == nil {
		emailFirstReq = cr.Email
		// 发送验证码存入session
		code := random.Code(4)
		session.Set("email_code", code)
		err := session.Save()
		if err != nil {
			global.Logger.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码是"+code)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		res.OkWithMessage("验证码已发送", c)
		return
	}
	code := session.Get("email_code")
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	// 修改邮箱
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if cr.Email != emailFirstReq {
		res.FailWithMessage("请求错误,非接受验证码邮箱", c)
		return
	}
	if err = pwd.CheckPasswordLever(cr.Password); err != nil {
		res.FailWithMessage(fmt.Sprint(err), c)
		return
	}
	hashPwd, _ := pwd.HashPwd(cr.Password)
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}
	res.OkWithMessage("邮箱更改成功", c)
}
