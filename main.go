package main

import (
	"gvb_server/core"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	// 读取配置文件
	core.InitCore()
	// log
	global.Logger = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// flag
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// router
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Logger.Infof("[gvb]  backend运行在 %s", addr)

	err := router.Run(global.Config.System.Addr())
	if err != nil {
		global.Logger.Error(err.Error())
	}
}
