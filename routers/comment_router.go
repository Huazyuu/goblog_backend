package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) CommentRouter() {
	commentApi := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), commentApi.CommentCreateView)
	router.GET("comments", commentApi.CommentListView)
	router.GET("comments/:id", commentApi.CommentDigg)
	router.DELETE("comments/:id", commentApi.CommentRemoveView)

}
