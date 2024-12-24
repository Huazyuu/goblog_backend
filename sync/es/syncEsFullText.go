package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/esServer"
)

func main() {
	core.InitCore("")
	core.InitLogger()
	global.ESClient = core.InitElasticSearch()

	boolSearch := elastic.NewMatchAllQuery()
	res, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).Do(context.Background())

	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)
		indexList := esServer.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)
		/*批量添加
		遍历 indexList，即需要添加到 Elasticsearch 的数据。
		elastic.NewBulkIndexRequest()每一条数据都会生成一个 BulkIndexRequest：
		.Index(models.FullTextModel{}.Index())：指定写入的索引（类似数据库中的表名）。
		.Doc(indexData)：指定要写入的数据内容（通常是 JSON 格式的文档）。
		bulk.Add(req)：将生成的 BulkIndexRequest 添加到 Bulk 请求中。*/
		bulk := global.ESClient.Bulk()
		for _, indexData := range indexList {
			req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
			bulk.Add(req)
		}
		result, err := bulk.Do(context.Background())
		if err != nil {
			logrus.Error(err)
			continue
		}
		fmt.Println(article.Title, "添加成功", "共", len(result.Succeeded()), " 条！")
	}

}
