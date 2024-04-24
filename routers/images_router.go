package routers

import "gvb_server/api"

func (router *RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi

	router.POST("images", imagesApi.ImageUploadView)   // 上传图片
	router.GET("images", imagesApi.ImageListView)      // 获取图片列表
	router.DELETE("images", imagesApi.ImageRemoveView) // 删除图片
	router.PUT("images", imagesApi.ImageUpdateView)    // 修改图片

}
