package cornServer

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redisServer"
)

// SyncArticleData 同步文章数据到es
func SyncArticleData() {
	// 1.查询es中的全部数据，为后面的数据更新做准备
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		return
	}
	// 2.拿到redis中的缓存数据
	diggInfo := redisServer.NewDigg().GetInfo()
	lookInfo := redisServer.NewArticleLook().GetInfo()
	commentInfo := redisServer.NewCommentCount().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		// 3. 计算新的数据   之前的数据+缓存中的数据
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment
		// 4. 要判断一下，是否有变化，如果三个变化之后的数据和之前的是一样的
		if digg == 0 && look == 0 && comment == 0 {
			global.Logger.Infof("%s无变化", article.Title)
			continue
		}
		// 5. 更新
		_, err = global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"look_count":    newLook,
				"comment_count": newComment,
				"digg_count":    newDigg,
			}).Do(context.Background())
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		global.Logger.Infof("%s 更新成功 点赞数：%d 评论数：%d  浏览量：%d",
			article.Title, newDigg, newComment, newLook)
	}
	// 6. 清除redis中的数据
	redisServer.NewDigg().Clear()
	redisServer.NewArticleLook().Clear()
	redisServer.NewCommentCount().Clear()
}
