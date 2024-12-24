package routers

import (
	"gvb_server/api"
)

func (router *RouterGroup) DiggRouter() {
	diggApi := api.ApiGroupApp.DiggApi
	router.POST("/digg/article", diggApi.DiggArticleView)
}
