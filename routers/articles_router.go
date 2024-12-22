package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) ArticlesRouter() {

	articlesApi := api.ApiGroupApp.ArticleApi
	// 发送文章
	router.POST("/articles", middleware.JwtAuth(), articlesApi.ArticleCreateView)
	// 查看文章list
	router.GET("/articles", articlesApi.ArticleListView)
	// 文章详细
	router.GET("/articles/detail", articlesApi.ArticleDetailByTitleView)
	// 文章发布时间列表
	router.GET("/articles/calendar", articlesApi.ArticleCalendarView)
	// 文章tag列表
	router.GET("/articles/tags", articlesApi.ArticleTagListView)
	// 更新文章
	router.PUT("/articles", middleware.JwtAuth(), articlesApi.ArticleUpdateView)
	// 删除文章(list)
	router.DELETE("/articles", middleware.JwtAuth(), articlesApi.ArticleRemoveView)
	// 文章id查询detail
	router.GET("/articles/:id", articlesApi.ArticleDetailView)
}
