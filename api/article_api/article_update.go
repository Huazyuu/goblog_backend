package article_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/esServer"
	"time"
)

type ArticleUpdateRequest struct {
	Title    string   `json:"title"`                 // 文章标题
	Abstract string   `json:"abstract"`              // 文章简介
	Content  string   `json:"content"`               // 文章内容
	Category string   `json:"category"`              // 文章分类
	Source   string   `json:"source"`                // 文章来源
	Link     string   `json:"link"`                  // 原文链接
	BannerID uint     `json:"banner_id"`             // 文章封面id
	Tags     []string `json:"tags"`                  // 文章标签
	ID       string   `json:"id" binding:"required"` // 需要修改的文章id
}

func (ArticlesApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	var bannerUrl string
	if cr.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).
			Where("id = ?", cr.BannerID).
			Select("path").Scan(&bannerUrl).
			Error
		if err != nil {
			res.FailWithMessage("banner 不存在", c)
			return
		}
	}
	article := models.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerUrl,
		Tags:      cr.Tags,
	}

	err = article.GetDataByID(cr.ID)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("文章不存在", c)
		return
	}

	// structs 转 maps
	maps := structs.Map(&article)
	var dataMap = make(map[string]any)
	// 去掉空值
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		dataMap[key] = v
	}

	err = esServer.ArticleUpdate(cr.ID, maps)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("文章更新失败", c)
		return
	}
	res.OkWithMessage("更新成功", c)
}
