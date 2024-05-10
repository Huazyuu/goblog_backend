package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gvb_server/api"
	"gvb_server/middleware"
)

var store = cookie.NewStore([]byte("ZYUUFORYUCOOKIESECRET"))

func (router *RouterGroup) UsersRouter() {
	usersApi := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))

	router.POST("users/email_login", usersApi.EmailLoginView)
	router.POST("users", usersApi.UserCreateView)
	// todo qq login router
	router.GET("login", usersApi.QQLoginView)

	// auth
	{
		router.POST("users/logout", middleware.JwtAuth(), usersApi.LogoutView)
		router.POST("users/user_bind_email", middleware.JwtAuth(), usersApi.UserBindEmailView)
		router.GET("users", middleware.JwtAuth(), usersApi.UserListView)
		router.PUT("users/users_pwd", middleware.JwtAuth(), usersApi.UserUpdatePasswordView)
	}

	// admin
	{
		router.PUT("users/users_role", middleware.JwtAdmin(), usersApi.UserUpdateRoleView)
		router.DELETE("users", middleware.JwtAdmin(), usersApi.UserRemoveView)
	}

}
