package advert_api

type AdvertApi struct {
}

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`       // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法"`   // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法"` // 图片
	IsShow bool   `json:"is_show" `                                   // 是否展示
}
