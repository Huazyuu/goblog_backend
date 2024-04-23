package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type Page struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
	Sort  string `form:"sort"`
}

// ImageListView ImageListView分页
func (imagesApi *ImagesApi) ImageListView(c *gin.Context) {
	var cr Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	// 分页
	cnt := global.DB.Find(&imageList).Select("id").RowsAffected
	// fmt.Println(cnt)
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}
	global.DB.Limit(cr.Limit).Offset(offset).Find(&imageList)

	res.OkWithData(gin.H{
		"count": cnt,
		"list":  imageList,
	}, c)
}
