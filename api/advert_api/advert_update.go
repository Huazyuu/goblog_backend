package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advert models.AdvertModel
	err = global.DB.Select("id").Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}
	//  Updates 方法默认会忽略零值
	// 使用 struct 更新时, GORM 将只更新非零值字段。 你可能想用 map 来更新属性，或者使用 Select 声明字段来更新
	/*err = global.DB.Model(&advert).Select("title", "href", "images", "is_show").Updates(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error*/
	//  将 is_show 的值更新为 false，Update 方法实际上不会执行任何操作，因为它将 false 视为零值。
	/*	err = global.DB.Model(&advert).
		Update("Title", cr.Title).
		Update("Href", cr.Href).
		Update("Images", cr.Images).
		Update("IsShow", cr.IsShow).Error*/
	// map 不会忽略 0 值 structs.Map()第三方库 结构体加`struct:""`tag
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}

	res.OkWithMessage("修改广告成功", c)
}
