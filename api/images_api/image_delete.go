package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// ImageRemoveView 删除图片
func (imagesApi *ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	// SELECT * FROM `banner_models` WHERE `banner_models`.`id` = cr.IDList
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	// DELETE FROM `banner_models` WHERE `banner_models`.`id` = 16
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), c)

}
