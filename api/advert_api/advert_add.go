package advert_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// AdvertCreatView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertCreatView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 判重
	var advert models.AdvertModel
	err = global.DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("图片已存在", c)
		return
	}

	// 入库
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)

}
