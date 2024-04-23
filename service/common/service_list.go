package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// ComList ComList用于显示分页(可排序)
func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	// model 用来推断 T 的 类型
	DB := global.DB
	// console mysql log
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MySqlLog})
	}
	// sort
	if option.Sort != "order" {
		option.Sort = "created_at desc" // 时间倒序
	} else {
		option.Sort = "created_at asc"
	}
	// cnt
	count = DB.Select("id").Find(&list).RowsAffected
	// 分页
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
