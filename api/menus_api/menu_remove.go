package menus_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuRemoveView 删除菜单
// @Tags 菜单管理
// @Summary 删除菜单
// @Description 删除菜单
// @Param data body models.RemoveRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/menus [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (MenusApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var menuList []models.MenuModel
	cnt := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if cnt == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// transactions
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 menu_banner
		// DELETE FROM `menu_banner_models` WHERE `menu_ban   ner_models`.`menu_id` IN (2,3)
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Logger.Error(err)
			return err
		}
		// 删除menu
		// DELETE FROM `menu_models` WHERE `menu_models`.`id` IN (2,3)
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Logger.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个菜单", cnt), c)
}
