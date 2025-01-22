package routers

import (
	"gvb_server/api"
)

func (router *RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	router.GET("tags", tagApi.TagListView)
	router.GET("tag_names", tagApi.TagNameListView)
	// todo tag router jwt middleware
	router.POST("tags", tagApi.TagCreateView)
	router.PUT("tags/:id", tagApi.TagUpdateView)
	router.DELETE("tags", tagApi.TagRemoveView)
}
