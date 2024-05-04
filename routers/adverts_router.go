package routers

import "gvb_server/api"

func (router *RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi

	router.POST("adverts", advertApi.AdvertCreatView)
	router.GET("adverts", advertApi.AdvertListView)
	router.PUT("adverts/:id", advertApi.AdvertUpdateView)
	router.DELETE("adverts", advertApi.AdvertRemoveView)

}
