package routers

import (
	"gvb_server/api"
)

func (router *RouterGroup) ChatRouter() {
	chatApi := api.ApiGroupApp.ChatApi
	router.GET("chat_groups", chatApi.ChatGroupView)
	router.GET("chat_groups_records", chatApi.ChatListView)
	router.DELETE("chat_groups_delete", chatApi.ChatRemoveView)
}
