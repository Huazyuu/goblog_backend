package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
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
	// swag
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// cors problem
	// router.Use(middleware.Cors()) // 解决跨域问题 自写中间件
	router.Use(cors.Default()) // 解决跨域问题 gin官方包 "github.com/gin-contrib/cors"

	// api group
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}

	// 系统配置api
	routerGroupApp.SettingsRouter()
	// images api
	routerGroupApp.ImagesRouter()
	// advert api
	routerGroupApp.AdvertRouter()
	// menu api
	routerGroupApp.MenuRouter()

	return router
}
