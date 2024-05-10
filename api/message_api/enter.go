package message_api

type MessageApi struct{}

type MessageRequest struct {
	SenderId   uint   `json:"sender_id" binding:"required"`
	ReceiverId uint   `json:"receiver_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}
