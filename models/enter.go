package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"Updated-at"`
}
