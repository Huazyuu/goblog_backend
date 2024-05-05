package flag

import (
	sysFlag "flag"
)

type Option struct {
	DB   bool
	User string
}

// Parse 解析命令行参数
func Parse() Option {
	db := sysFlag.Bool("db", false, "初始化数据库 -db")
	user := sysFlag.String("u", "", "创建用户 -u user")
	sysFlag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB || option.User == "admin" || option.User == "user" {
		return true
	}
	return false
}

// SwitchOption 根据命令执行不同函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigration()
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
	} else {
		sysFlag.Usage()
	}

}
