package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/article_api"
	"gvb_server/api/chat_api"
	"gvb_server/api/comment_api"
	"gvb_server/api/data_api"
	"gvb_server/api/images_api"
	"gvb_server/api/log_api"
	"gvb_server/api/menus_api"
	"gvb_server/api/message_api"
	"gvb_server/api/settings_api"
	"gvb_server/api/tag_api"
	"gvb_server/api/users_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menus_api.MenusApi
	UserApi     users_api.UsersApi
	TagApi      tag_api.TagApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticlesApi
	CommentApi  comment_api.CommentApi
	ChatApi     chat_api.ChatApi
	LogApi      log_api.LogApi
	DateApi     data_api.DataApi
}

var ApiGroupApp = new(ApiGroup)
