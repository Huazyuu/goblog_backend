package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}
type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenusApi) MenuListView(c *gin.Context) {
	// menu
	var menuList []models.MenuModel
	var menuIDList []uint
	//  SELECT * FROM `menu_models` ORDER BY sort desc
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	// menuBanner
	var menuBanners []models.MenuBannerModel

	// SELECT `id` FROM `menu_models` ORDER BY sort desc
	//  SELECT * FROM `banner_models` WHERE `banner_models`.`id` IN (3,2,1)
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	// resp
	var menus []MenuResponse
	for _, model := range menuList {
		// model就是一个菜单
		var banners []Banner
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
	return
}
