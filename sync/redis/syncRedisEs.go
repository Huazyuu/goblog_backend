package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redisServer"
)

func init() {
	core.InitCore("settings.yaml")
	global.Redis = core.InitRedis()
	global.ESClient = core.InitElasticSearch()
	global.Logger = core.InitLogger()
}
func main() {
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())

	if err != nil {
		logrus.Error(err)
		return
	}

	diggInfo := redisServer.NewDigg().GetInfo()
	lookInfo := redisServer.NewArticleLook().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)

		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look

		if article.DiggCount == newDigg && article.LookCount == newLook {
			global.Logger.Warn(article.Title, " 点赞数和浏览量无变化, 点赞数", article.DiggCount, "\t浏览量", article.LookCount)
			continue
		}

		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count": newDigg,
				"look_count": newLook,
			}).
			Do(context.Background())
		if err != nil {
			global.Logger.Error(err.Error())
			continue
		}
		global.Logger.Info(article.Title, "点赞数浏览量同步成功, 点赞数", newDigg, "\t浏览量", article.LookCount)
	}
	redisServer.NewDigg().Clear()
	redisServer.NewArticleLook().Clear()
}
