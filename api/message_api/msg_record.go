package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户id"`
}

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
		res.FailWithMessage("查询与对方的不可为自己的id", c)
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
