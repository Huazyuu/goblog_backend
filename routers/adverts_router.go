package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi
	router.GET("adverts", advertApi.AdvertListView)

	router.POST("adverts", middleware.JwtAdmin(), advertApi.AdvertCreateView)
	router.PUT("adverts/:id", middleware.JwtAdmin(), advertApi.AdvertUpdateView)
	router.DELETE("adverts", middleware.JwtAdmin(), advertApi.AdvertRemoveView)

}
