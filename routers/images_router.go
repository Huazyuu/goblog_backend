package routers

import "gvb_server/api"

func (router *RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi

	router.POST("images", imagesApi.ImageUploadView)
	router.GET("images", imagesApi.ImageListView)
	router.DELETE("images", imagesApi.ImageRemoveView)

}
