package routers

import "gvb_server/api"

func (router *RouterGroup) MenusRouter() {
	MenuApi := api.ApiGroupApp.MenuApi

	router.POST("menus", MenuApi.MenuCreateView)

	router.GET("menus", MenuApi.MenuListView)
	router.GET("menu_names", MenuApi.MenuNameList)
	router.GET("menus/:id", MenuApi.MenuDetailView)

	router.DELETE("menus", MenuApi.MenuRemoveView)

	router.PUT("menus/:id", MenuApi.MenuUpdateView)

}
