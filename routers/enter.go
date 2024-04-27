package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	// gin.DisableConsoleColor()
	// file, _ := os.Open("./log/myLog.log")
	// gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}

	// 系统配置api
	routerGroupApp.SettingsRouter()
	// images api
	routerGroupApp.ImagesRouter()
	// advert api
	routerGroupApp.AdvertRouter()

	return router
}
