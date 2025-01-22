package users_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
	"gvb_server/utils"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
)

type EmailLoginRequest struct {
	Username string `json:"username" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UsersApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	log := log_stash.NewLogGin(c)
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.Username, cr.Username).Error
	if err != nil {
		// 没找到
		global.Logger.Warn("用户名不存在")
		log.Warn(fmt.Sprintf("%s 用户名不存在", cr.Username))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// check pwd
	isCheck, _ := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Logger.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("用户名密码错误%s %s", cr.Username, cr.Password))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("token生成失败", c)
		log.Error(fmt.Sprintf("token生成失败%s", err.Error()))
		return
	}

	ip, addr := utils.GetAddrByGin(c)
	log = log_stash.New(ip, token)
	log.Info("登陆成功")
	// 入库
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		NickName:  userModel.NickName,
		IP:        ip,
		Addr:      addr,
		Token:     token,
		Device:    "",
		LoginType: ctype.SignEmail,
	})
	res.OkWithData(token, c)
}
