package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MessageCreateView 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var sender, receiver models.UserModel
	if err = global.DB.Take(&sender, cr.SenderId).Error; err != nil {
		res.FailWithMessage("发送人不存在", c)
		return
	}
	if err = global.DB.Take(&receiver, cr.ReceiverId).Error; err != nil {
		res.FailWithMessage("接收人不存在", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SenderId,
		SendUserNickName: sender.NickName,
		SendUserAvatar:   sender.Avatar,
		RevUserID:        cr.ReceiverId,
		RevUserNickName:  receiver.NickName,
		RevUserAvatar:    receiver.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		res.FailWithMessage("消息发送失败", c)
		return
	}
	res.OkWithMessage("消息发送成功", c)
}
