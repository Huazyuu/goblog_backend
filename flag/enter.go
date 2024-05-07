package flag

import (
	sysFlag "flag"
	"github.com/fatih/structs"
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
func IsWebStop(option Option) (flag bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				flag = true
			}
		case bool:
			if val {
				flag = true
			}
		}
	}
	return
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
