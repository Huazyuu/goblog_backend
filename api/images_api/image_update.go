package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImageUpdateRequest struct {
	ID   int    `json:"id" binding:"required" msg:"请输入文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

func (imagesApi *ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imageModel models.BannerModel
	// Take 获取一条记录
	if err = global.DB.Take(&imageModel, cr.ID).Error; err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	if err = global.DB.Model(&imageModel).Update("name", cr.Name).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
	return

}
