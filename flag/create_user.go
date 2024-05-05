package flag

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/pwd"
)

func CreateUser(per string) {
	var (
		nickName   string
		userName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名: ")
	if cnt, _ := fmt.Scan(&userName); cnt == 0 {
		global.Logger.Error("用户名不许为空")
		return
	}
	fmt.Printf("请输入昵称: ")
	if cnt, _ := fmt.Scan(&nickName); cnt == 0 {
		global.Logger.Error("昵称不许为空")
		return
	}
	fmt.Printf("请输入密码: ")
	if cnt, _ := fmt.Scan(&password); cnt == 0 {
		global.Logger.Error("密码不许为空")
		return
	}
	fmt.Printf("请再次输入密码: ")
	_, _ = fmt.Scan(&rePassword)
	fmt.Printf("请输入邮箱(可选): ")
	_, _ = fmt.Scan(&email)

	// userName check
	var userModel models.UserModel
	err := global.DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Logger.Error("用户名已存在,请重新输入")
		return
	}
	// pwd check
	if password != rePassword {
		global.Logger.Error("密码不一致,请重新输入")
		return
	}
	// hash pwd
	hashPwd, _ := pwd.HashPwd(password)
	// permission
	role := ctype.PermissionUser
	if per == "admin" {
		role = ctype.PermissionAdmin
	}
	// avatar
	avatar := "/uploads/avatar/avatar_default.png"
	// db insert
	err = global.DB.Create(&models.UserModel{
		NickName:       nickName,
		UserName:       userName,
		Password:       hashPwd,
		Avatar:         avatar,
		Email:          email,
		Tel:            "",
		IP:             "127.0.0.1",
		Addr:           "内网地址",
		Token:          "",
		Role:           role,
		SignStatus:     ctype.SignEmail,
		ArticleModels:  nil,
		CollectsModels: nil,
	}).Error
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Infof("创建用户%s成功", userName)

}
