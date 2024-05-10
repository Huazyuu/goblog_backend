package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/service/userServer"
)

func CreateUser(per string) {
	// input
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
	fmt.Printf("请输入邮箱: ")
	_, _ = fmt.Scan(&email)
	fmt.Printf("请输入密码: ")
	if cnt, _ := fmt.Scan(&password); cnt == 0 {
		global.Logger.Error("密码不许为空")
		return
	}
	fmt.Printf("请再次输入密码: ")
	_, _ = fmt.Scan(&rePassword)

	// pwd check
	if password != rePassword {
		global.Logger.Error("密码不一致,请重新输入")
		return
	}

	// permission
	role := ctype.PermissionUser
	if per == "admin" {
		role = ctype.PermissionAdmin
	}

	err := userServer.UserService{}.CreateUser(userName, nickName, password, role, email, "127.0.0.1")
	if err != nil {
		global.Logger.Error(err)
		return
	}

	global.Logger.Infof("创建用户%s成功", userName)

}
