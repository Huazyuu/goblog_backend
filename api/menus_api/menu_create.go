package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`                          // 简介,可由多个,类型为[]string
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                // 简介切换的时间，单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    // 图篇切换的时间，单位秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` // 菜单的顺序
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          // 具体图片的顺序
}

// MenuCreateView 发布菜
// MenuCreateView 发布菜单
// @Tags 菜单管理
// @Summary 发布菜单
// @Description 发布菜单
// @Param data body MenuRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/menus [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (MenusApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// MenuModel 入库
	// 查重
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("重复的菜单", c)
		return
	}
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("添加菜单失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}
	// MenuBannerModel 入库
	var menuBannerList []models.MenuBannerModel
	for _, v := range cr.ImageSortList {
		var imgModel models.BannerModel
		cnt := global.DB.Find(&imgModel, "id = ?", v.ImageID).RowsAffected
		if cnt <= 0 {
			res.FailWithMessage("没有这张图片", c)
			return
		}
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: v.ImageID,
			Sort:     v.Sort,
		})
	}
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}
