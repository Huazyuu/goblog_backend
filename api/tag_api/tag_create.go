package tag_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 判重
	var tag models.TagModel
	err = global.DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("标签已存在", c)
		return
	}

	// 入库
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		res.FailWithMessage("添加标签失败", c)
		return
	}
	res.OkWithMessage("添加标签成功", c)

}
