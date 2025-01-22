package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
	"gvb_server/utils"
)

// swagger api_doc
// http://127.0.0.1:8080/swagger/index.html#/

// @title gvb-server API文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitCore("")
	// log
	global.Logger = core.InitLogger()
	// gorm
	global.DB = core.InitGorm()

	// addr ip
	global.AddrDB = core.InitAddrDB()
	defer global.AddrDB.Close()

	// flag
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	// redis
	global.Redis = core.InitRedis()
	// es
	global.ESClient = core.InitElasticSearch()

	// router
	router := routers.InitRouter()
	// run
	utils.PrintSystemInfo()
	err := router.Run(global.Config.System.Addr())
	if err != nil {
		global.Logger.Error(err.Error())
	}
}
