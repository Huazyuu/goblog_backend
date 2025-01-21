package routers

import "gvb_server/api"

func (router *RouterGroup) DataRouter() {
	dataApi := api.ApiGroupApp.DateApi
	router.GET("/date_login", dataApi.SevenLoginView)
	router.GET("/data_sum", dataApi.DataSumView)
}
