package chat_api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
	randomname "gvb_server/utils/random_name"
	"net/http"
	"strings"
	"time"
)

// ChatGroupView 群聊
func (ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权 true放行 false拦截
			return true
		},
	}
	// http 升级 websockets protocol
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// add user
	addr := conn.RemoteAddr().String()
	nickname := randomname.GenerateName()
	avatar := fmt.Sprintf("uploads/chat_avatar/%s.png", string([]rune(nickname)[0]))
	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickname,
		Avatar:   avatar,
	}
	ConnGroupMap[addr] = chatUser
	global.Logger.Infof("[%s][%s]连接成功", addr, nickname)

	// 数据处理
	for {
		// 读取聊天数据 json格式
		_, content, err := conn.ReadMessage()
		if err != nil {
			// 断开连接
			sendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     fmt.Sprintf("[%s]离开聊天室", chatUser.NickName),
				Date:        time.Now(),
				MsgType:     OutRoomMsg,
				OnlineCount: len(ConnGroupMap) - 1,
			})
			break
		}
		// json数据序列化request结构体
		var req GroupRequest
		err = json.Unmarshal(content, &req)
		if err != nil {
			// 参数绑定失败
			sendMsg(addr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     SystemMsg,
				Content:     "参数绑定失败",
				OnlineCount: len(ConnGroupMap),
			})
			continue
		}
		// 断言
		switch req.MsgType {
		case TextMsg:
			// 内容为空
			if strings.TrimSpace(req.Content) == "" {
				sendMsg(addr, GroupResponse{
					NickName:    chatUser.NickName,
					Avatar:      chatUser.Avatar,
					MsgType:     SystemMsg,
					Content:     "消息不能为空",
					OnlineCount: len(ConnGroupMap),
				})
				continue
			}
			sendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     req.Content,
				MsgType:     TextMsg,
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap),
			})
		case InRoomMsg:
			sendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     fmt.Sprintf("[%s]进入聊天室", chatUser.NickName),
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap),
			})
		default:
			sendMsg(addr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     SystemMsg,
				Content:     "消息类型错误",
				OnlineCount: len(ConnGroupMap),
			})
		}

	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// sendGroupMsg 群聊功能
func sendGroupMsg(conn *websocket.Conn, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	_addr := conn.RemoteAddr().String()
	ip, addr := getIPAndAddr(_addr)

	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  true,
		MsgType:  response.MsgType,
	})
	for _, chatUser := range ConnGroupMap {
		chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// sendMsg 给某个用户发消息
func sendMsg(_addr string, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	chatUser := ConnGroupMap[_addr]
	ip, addr := getIPAndAddr(_addr)
	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  false,
		MsgType:  response.MsgType,
	})
	chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

func getIPAndAddr(_addr string) (ip string, addr string) {
	addrList := strings.Split(_addr, ":")
	addr = utils.GetAddr(ip)
	return addrList[0], addr
}
