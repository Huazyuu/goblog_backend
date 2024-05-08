package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) MenusRouter() {
	MenuApi := api.ApiGroupApp.MenuApi

	router.POST("menus", middleware.JwtAdmin(), MenuApi.MenuCreateView)

	router.GET("menus", MenuApi.MenuListView)
	router.GET("menu_names", MenuApi.MenuNameList)
	router.GET("menus/:id", MenuApi.MenuDetailView)

	router.DELETE("menus", middleware.JwtAdmin(), MenuApi.MenuRemoveView)

	router.PUT("menus/:id", middleware.JwtAdmin(), MenuApi.MenuUpdateView)

}
