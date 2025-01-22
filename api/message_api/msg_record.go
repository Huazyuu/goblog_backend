package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
)

// MessageRecordView 用户的消息记录
// @Tags 消息管理
// @Summary 用户的消息记录
// @Description 用户的消息记录
// @Router /api/messages_record [post]
// @Param token header string  true  "token"
// @Param data body MessageRecordRequest  true  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=[]models.MessageModel}
func (MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	global.DB.Order("created_at").
		Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)

	if claims.UserID == cr.UserID {
		res.FailWithMessage("没有自身聊天消息,请查询其他id", c)
		return
	}

	for _, model := range _messageList {
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}

	// 点开消息，里面的每一条消息，都从未读变成已读

	res.OkWithData(messageList, c)
}
