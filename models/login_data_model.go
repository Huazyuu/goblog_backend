package models

import "gvb_server/models/ctype"

// LoginDataModel 统计用户登录数据 id, 用户id, 用户昵称，用户token，登录设备，登录时间
type LoginDataModel struct {
	MODEL

	IP        string           `gorm:"size:20" json:"ip"` // 登录的ip
	NickName  string           `gorm:"size:42" json:"nick_name"`
	Token     string           `gorm:"size:256" json:"token"`
	Device    string           `gorm:"size:256" json:"device"` // 登录设备
	Addr      string           `gorm:"size:64" json:"addr"`
	LoginType ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"` // 登录类型

	// 外键
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
}
