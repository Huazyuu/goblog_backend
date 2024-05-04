package models

import "gvb_server/models/ctype"

// MenuModel 菜单表
type MenuModel struct {
	MODEL
	Title        string      `gorm:"size:32" json:"title"` // 标题
	Path         string      `gorm:"size:32" json:"path"`
	Slogan       string      `gorm:"size:64" json:"slogan"`       // slogan
	Abstract     ctype.Array `gorm:"type:string" json:"abstract"` // 简介
	AbstractTime int         `json:"abstract_time"`               // 简介的切换时间
	BannerTime   int         `json:"banner_time"`                 // 菜单图片的切换时间 为 0 表示不切换
	Sort         int         `gorm:"size:10" json:"sort"`         // 菜单的顺序

	// 创建 menu_banner_models 连接表
	// menu_banner_models(menu_id) -> MenuModel (id)
	// menu_banner_models(banner_id) -> BannerModel (id)
	Banners []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"` // 菜单的图片列表
}
