package main

import (
	"gvb_server/core"
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
	// router
	router := routers.InitRouter()
	router.Run(global.Config.System.Addr())
}
