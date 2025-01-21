package routers

import "gvb_server/api"

func (router *RouterGroup) LogRouter() {
	logApi := api.ApiGroupApp.LogApi
	router.GET("logs", logApi.LogListView)
	router.DELETE("logs", logApi.LogRemoveListView)
}
