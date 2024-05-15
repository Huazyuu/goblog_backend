package message_api

import "time"

type MessageApi struct{}

type MessageRequest struct {
	SenderId   uint   `json:"sender_id" binding:"required"`
	ReceiverId uint   `json:"receiver_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type Message struct {
	SendUserID       uint   `json:"sender_id"` // 发送人id
	SendUserNickName string `json:"sender_nickname"`
	SendUserAvatar   string `json:"sender_avatar"`

	ReceiveUserID       uint   ` json:"receiver_id"` // 接收人id
	ReceiveUserNickName string ` json:"receiver_nickname"`
	ReceiveUserAvatar   string `json:"receiver_avatar"`

	Content        string    `json:"content"` // 消息内容
	CreateAt       time.Time `json:"create_at"`
	MessageContent int       `json:"message_content"`
}

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户id"`
}
