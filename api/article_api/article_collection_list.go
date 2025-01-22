package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/jwt"
)

type CollectionsResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

// ArticleCollectionListView 用户收藏的文章列表
// @Tags 文章管理
// @Summary 用户收藏的文章列表
// @Description 用户收藏的文章列表
// @Param data query models.PageInfo  true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles/collects [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[CollResponse]}
func (ArticlesApi) ArticleCollectionListView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var articleIDList []interface{}
	list, cnt, err := common.ComList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
	})

	collMap := make(map[string]string)
	// fmt.Println(list)
	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}

	// term 完全匹配不会分词
	// term类似于MySQL的 where province=？
	// terms类似于MySQL中的 where province in (?, ? ,?)
	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)

	var collList = make([]CollectionsResponse, 0)

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	// fmt.Println(result.Hits.TotalHits.Value, articleIDList)

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		article.ID = hit.Id
		collList = append(collList, CollectionsResponse{
			ArticleModel: article,
			CreatedAt:    collMap[hit.Id],
		})
	}
	res.OkWithList(collList, cnt, c)
}
