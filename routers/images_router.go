package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router *RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi

	router.GET("images", imagesApi.ImageListView)                             // 获取图片列表
	router.POST("images", middleware.JwtAuth(), imagesApi.ImageUploadView)    // 上传图片
	router.DELETE("images", middleware.JwtAdmin(), imagesApi.ImageRemoveView) // 删除图片
	router.PUT("images", middleware.JwtAdmin(), imagesApi.ImageUpdateView)    // 修改图片

}
