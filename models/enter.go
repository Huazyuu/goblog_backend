package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"Updated-at"`
}

// PageInfo 分页Page
type PageInfo struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
	Sort  string `form:"sort"`
}
