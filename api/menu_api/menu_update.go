package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")

	// 清空banner
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	// 如果选择banner就添加
	if len(cr.ImageSortList) > 0 {
		var bannerList []models.MenuBannerModel
		for _, imgSort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				Sort:     imgSort.Sort,
				MenuID:   menuModel.ID,
				BannerID: imgSort.ImageID,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Logger.Error(err)
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}
	// 没选择banner
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}

	res.OkWithMessage("修改菜单成功", c)
}
