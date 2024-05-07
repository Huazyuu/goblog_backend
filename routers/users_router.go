package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) UsersRouter() {
	usersApi := api.ApiGroupApp.UserApi
	router.POST("users/email_login", usersApi.EmailLoginView)
	router.Use(middleware.JwtAuth())
	router.GET("users", usersApi.UserListView)
}