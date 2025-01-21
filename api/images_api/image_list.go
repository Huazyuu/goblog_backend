package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// ImageListView 图片列表查询(支持分页,排序) // todo 模糊搜索没做
func (imagesApi *ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo

	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	imageList, cnt, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})

	res.OkWithList(imageList, cnt, c)
}
