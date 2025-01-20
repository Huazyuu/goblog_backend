package chat_api

import (
	"github.com/gorilla/websocket"
	"gvb_server/models/ctype"
	"time"
)

type ChatApi struct{}

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"` // 前端自己生成
	Avatar   string `json:"avatar"`    // 头像
}

// ConnGroupMap 用户map
var ConnGroupMap = make(map[string]ChatUser)

const (
	InRoomMsg ctype.MsgType = iota + 1
	TextMsg
	ImageMsg
	VoiceMsg
	VideoMsg
	SystemMsg
	OutRoomMsg
)

// GroupRequest 请求
type GroupRequest struct {
	Content string        `json:"content"`  // 聊天的内容
	MsgType ctype.MsgType `json:"msg_type"` // 聊天类型
}

// GroupResponse 响应
type GroupResponse struct {
	NickName    string        `json:"nick_name"`
	Avatar      string        `json:"avatar"`       // 头像
	MsgType     ctype.MsgType `json:"msg_type"`     // 聊天类型
	Content     string        `json:"content"`      // 聊天的内容
	Date        time.Time     `json:"date"`         // 消息的时间
	OnlineCount int           `json:"online_count"` // 在线人数
}
