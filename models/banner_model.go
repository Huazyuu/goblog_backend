package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

// BannerModel img表
type BannerModel struct {
	MODEL
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片的类型， 本地还是七牛
}

// BeforeDelete hook
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Logger.Error(err)
			return err
		}
	}
	return nil
}
