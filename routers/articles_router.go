package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) ArticlesRouter() {

	articlesApi := api.ApiGroupApp.ArticleApi
	router.POST("/articles", middleware.JwtAuth(), articlesApi.ArticleCreateView)
	router.GET("/articles", articlesApi.ArticleListView)
	router.GET("/articles/detail", articlesApi.ArticleDetailByTitleView)
	router.GET("/articles/calendar", articlesApi.ArticleCalendarView)
	router.GET("/articles/:id", articlesApi.ArticleDetailView)
}
