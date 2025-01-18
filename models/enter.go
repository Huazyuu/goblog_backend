package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey,select($any)" json:"id"`
	CreatedAt time.Time `json:"created-at,select($any)"`
	UpdatedAt time.Time `json:"updated-at,select($any)"`
}

// PageInfo 分页Page
type PageInfo struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
	Sort  string `form:"sort"`
}

// RemoveRequest 删除请求
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
