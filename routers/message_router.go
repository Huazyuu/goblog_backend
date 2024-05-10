package routers

import "gvb_server/api"

func (router *RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	router.POST("messages", messageApi.MessageCreateView)
}
