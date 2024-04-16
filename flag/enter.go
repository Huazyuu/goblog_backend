package flag

import sysFlag "flag"

type Option struct {
	DB bool
}

// Parse 解析命令行参数
func Parse() Option {
	db := sysFlag.Bool("db", false, "初始化数据库")
	sysFlag.Parse()
	return Option{
		DB: *db,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	return option.DB
}

// SwitchOption 根据命令执行不同函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigration()
	}
}
