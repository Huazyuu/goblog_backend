package main

import (
	"fmt"
	"github.com/fatih/structs"
	"gvb_server/models"
)

// struct 转 map

type AdvertRequest struct {
	models.MODEL `structs:"-"` // "-" 忽略
	Title        string        `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        // 显示的标题
	Href         string        `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`     // 跳转链接
	Images       string        `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` // 图片
	IsShow       bool          `json:"is_show" structs:"is_show"`                                   // 是否展示
}

func main() {
	u := AdvertRequest{
		Title:  "t",
		Href:   "h",
		Images: "i",
		IsShow: false,
	}
	m := structs.Map(&u)
	fmt.Println(m)
}
