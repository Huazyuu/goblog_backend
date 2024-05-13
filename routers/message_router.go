package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	// user
	router.POST("messages", middleware.JwtAuth(), messageApi.MessageCreateView)
	router.GET("messages", middleware.JwtAuth(), messageApi.MessageListView)
	router.GET("messages_record", middleware.JwtAuth(), messageApi.MessageRecordView)
	// admin
	router.GET("messages_all", middleware.JwtAdmin(), messageApi.MessageListAllView)

}
