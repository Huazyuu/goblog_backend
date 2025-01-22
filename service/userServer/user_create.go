package userServer

import (
	"errors"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
	"gvb_server/utils/pwd"
)

var avatar = "/uploads/avatar/avatar_default.png"

func (UserService) CreateUser(username, nickname, password string, role ctype.Role, email, ip string) error {
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name= ? ", username).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	// hash pwd
	hashPwd, _ := pwd.HashPwd(password)
	err = global.DB.Create(&models.UserModel{
		NickName:   nickname,
		UserName:   username,
		Password:   hashPwd,
		Avatar:     avatar,
		Email:      email,
		IP:         ip,
		Addr:       utils.GetAddr(ip),
		Role:       role,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
