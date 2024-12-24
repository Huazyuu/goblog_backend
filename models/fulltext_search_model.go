package models

import (
	"context"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"os"
)

// FullTextModel 全文搜索
type FullTextModel struct {
	// structs struct 转 map 后的key
	ID    string `structs:"id" json:"id"`       // es的id
	Title string `structs:"title" json:"title"` // 文章标题
	Slug  string `structs:"slug" json:"slug"`   // 包含文章的id 的跳转地址
	Body  string `structs:"body" json:"body"`   // 文章内容
}

func (FullTextModel) Index() string {
	return "full_text_index"
}

func (FullTextModel) Mapping() string {
	path := "models/fulltext_mapper.json"
	txt, err := os.ReadFile(path)
	if err != nil {
		logrus.Error(err)
		return err.Error()
	}
	return string(txt)
}

// IndexExists 索引是否存在
func (a FullTextModel) IndexExists() bool {
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
func (a FullTextModel) CreateIndex() error {
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
func (a FullTextModel) RemoveIndex() error {
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
