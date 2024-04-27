package routers

import "gvb_server/api"

func (router *RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi

	router.POST("adverts", advertApi.AdvertCreatView)
	// router.GET("advert")
	// router.PUT("advert")
	// router.DELETE("advert")

}
