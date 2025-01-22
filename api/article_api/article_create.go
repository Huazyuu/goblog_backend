package article_api

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/esServer"
	"gvb_server/utils/jwt"
	"math/rand"
	"strings"
	"time"
)

type ArticleRequest struct {
	Title string `json:"title" binding:"required" msg:"文章标题必填"` // 文章标题
	// Keyword string      `json:"keywords" msg:"Keywords"`
	Abstract string      `json:"abstract"`                                // 文章简介
	Content  string      `json:"content" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string      `json:"category"`                                // 文章分类
	Source   string      `json:"source"`                                  // 文章来源
	Link     string      `json:"link"`                                    // 原文链接
	BannerID uint        `json:"banner_id"`                               // 文章封面id
	Tags     ctype.Array `json:"tags"`                                    // 文章标签
}

// ArticleCreateView 发布文章
// @Tags 文章管理
// @Summary 发布文章
// @Description 发布文章
// @Param data body ArticleRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticlesApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	userID := claims.UserID
	userNickName := claims.NickName

	// 处理content md->html
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// html 获取内容
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	nodes := doc.Find("script").Nodes
	// 过滤<script> 防止xss
	if len(nodes) > 0 {
		doc.Find("script").Remove()
		// html -> md
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}

	// abstract
	if cr.Abstract == "" {
		// abs为空取content前100字符,由于中英文字节数不同,使用rune存取
		abs := []rune(doc.Text())
		// content 转化 html 过滤掉xss
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs)
		}
	}

	// 不传banner随机图片
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(&models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			res.FailWithMessage("没有banner数据", c)
			return
		}
		cr.BannerID = bannerIDList[rand.Intn(len(bannerIDList))]
	}

	// banner_id 的 url
	var bannerUrl string
	err = global.DB.Model(&models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		res.FailWithMessage("banner图片不存在", c)
		return
	}

	// avatar 用户头像
	var avatar string
	err = global.DB.Model(&models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		res.FailWithMessage("avatar头像不存在", c)
		return
	}

	article := models.ArticleModel{
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	// 标题去重
	if article.ISExistTitle() {
		res.FailWithMessage("文章已存在", c)
		return
	}

	err = article.Create()
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	go esServer.ASyncArticleFullText(article.ID, article.Title, article.Content)
	res.OkWithMessage("文章发布成功", c)
}
