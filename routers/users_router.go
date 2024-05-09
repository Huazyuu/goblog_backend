package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) UsersRouter() {
	usersApi := api.ApiGroupApp.UserApi
	router.POST("users/email_login", usersApi.EmailLoginView)
	router.POST("users/logout", middleware.JwtAuth(), usersApi.LogoutView)

	router.GET("users", middleware.JwtAuth(), usersApi.UserListView)

	router.PUT("users/users_pwd", middleware.JwtAuth(), usersApi.UserUpdatePasswordView)
	router.PUT("users/users_role", middleware.JwtAdmin(), usersApi.UserUpdateRoleView)

	router.DELETE("users", middleware.JwtAdmin(), usersApi.UserRemoveView)

}
