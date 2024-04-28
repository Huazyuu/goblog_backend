package routers

import "gvb_server/api"

func (router *RouterGroup) MenuRouter() {
	MenuApi := api.ApiGroupApp.MenuApi

	router.POST("menus", MenuApi.MenuCreateView)
	router.GET("menus", MenuApi.MenuListView)
	router.DELETE("menus", MenuApi.MenuRemoveView)
	router.PUT("menus", MenuApi.MenuUpdateView)

}
