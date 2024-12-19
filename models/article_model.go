package models

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models/ctype"
)

// ArticleModel 文章表 es存取 index相当于表名 mapping表结构
type ArticleModel struct {
	// structs struct 转 map 后的key
	ID        string `structs:"id" json:"id"`                 // es的id
	CreatedAt string `structs:"created_at" json:"created_at"` // 创建时间
	UpdatedAt string `structs:"updated_at" json:"updated_at"` // 更新时间

	Title    string `structs:"title" json:"title"`                // 文章标题
	Keyword  string `structs:"keyword" json:"keyword,omit(list)"` // 关键字
	Abstract string `structs:"abstract" json:"abstract"`          // 文章简介
	Content  string `structs:"content" json:"content,omit(list)"` // 文章内容

	LookCount     int `structs:"look_count" json:"look_count"`         // 浏览量
	CommentCount  int `structs:"comment_count" json:"comment_count"`   // 评论量
	DiggCount     int `structs:"digg_count" json:"digg_count"`         // 点赞量
	CollectsCount int `structs:"collects_count" json:"collects_count"` // 收藏量

	UserID       uint   `structs:"user_id" json:"user_id"`               // 用户id
	UserNickName string `structs:"user_nick_name" json:"user_nick_name"` // 用户昵称
	UserAvatar   string `json:"user_avatar"`                             // 用户头像

	Category string `structs:"category" json:"category"`        // 文章分类
	Source   string `structs:"source" json:"source,omit(list)"` // 文章来源
	Link     string `structs:"link" json:"link,omit(list)"`     // 原文链接

	BannerID  uint   `structs:"banner_id" json:"banner_id"`   // 文章封面id
	BannerUrl string `structs:"banner_url" json:"banner_url"` // 文章封面

	Tags ctype.Array `structs:"tags" json:"tags"` // 文章标签
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "source": { 
        "type": "keyword"
      },
      "link": { 
        "type": "keyword"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "keyword"
      }, 
      "tags": { 
        "type": "keyword"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "updated_at":{
        "type": "date",
        "null_value": "null",
        "format": "yyyy-MM-dd HH:mm:ss"
      }
    }
  }
}
`
}

// IndexExists 索引是否存在
func (a ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		_ = a.RemoveIndex()
	}
	// 没有索引 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", a.Index())
	return nil
}

// RemoveIndex 删除索引
func (a ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}

// Create 添加的方法
func (a ArticleModel) Create() (err error) {
	indexResponse, err := global.ESClient.Index().
		Index(a.Index()).
		BodyJson(a).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// ISExistTitle 是否存在该文章
func (a ArticleModel) ISExistTitle() bool {
	res, err := global.ESClient.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	// fmt.Println(res.Hits.TotalHits.Value)
	if err != nil {
		logrus.Error(err.Error())
		return false
	}

	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}
