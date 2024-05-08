package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi

	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoUpdateView)

}
