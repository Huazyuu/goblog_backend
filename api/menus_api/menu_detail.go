package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuDetailView 菜单详情
// @Tags 菜单管理
// @Summary 菜单详情
// @Description 菜单详情
// @Param id path int  true  "id"
// @Router /api/menus/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=MenuResponse}
func (MenusApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// menuBanner
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Find(&menuBanners, "menu_id=?", id)

	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
	return
}
