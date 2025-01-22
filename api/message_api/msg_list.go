package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
)

type MessageGroup map[uint]*Message

// MessageListView 用户与其他人的消息列表
// @Tags 消息管理
// @Summary 用户与其他人的消息列表
// @Description 用户与其他人的消息列表
// @Router /api/messages [get]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{data=[]Message}
func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var (
		msgGroup = MessageGroup{}
		msgList  []models.MessageModel // gorm handle list
		msgs     []Message             // response list
	)

	global.DB.Order("created_at").
		Find(&msgList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range msgList {
		msg := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,

			ReceiveUserID:       model.RevUserID,
			ReceiveUserNickName: model.RevUserNickName,
			ReceiveUserAvatar:   model.RevUserAvatar,

			Content:        model.Content,
			CreateAt:       model.CreatedAt,
			MessageContent: 1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := msgGroup[idNum]
		if !ok {
			// 不存在
			msgGroup[idNum] = &msg
			continue
		}
		msg.MessageContent = val.MessageContent + 1
		msgGroup[idNum] = &msg
	}
	for _, msg := range msgGroup {
		msgs = append(msgs, *msg)
	}
	res.OkWithData(msgs, c)
}
