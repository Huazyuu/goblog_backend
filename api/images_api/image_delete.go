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
	/*
		可以将一个主键切片传递给Delete 方法，以便更高效的删除数据量大的记录
		var users = []User{{ID: 1}, {ID: 2}, {ID: 3}}
		db.Delete(&users)
		// DELETE FROM users WHERE id IN (1,2,3);
	*/
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
