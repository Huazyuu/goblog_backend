package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/esServer"
	"gvb_server/utils/jwt"
)

type ESIDListRequest struct {
	IDList []string `json:"id_list"`
}

// ArticleCollBatchRemoveView 用户取消收藏文章
// @Tags 文章管理
// @Summary 用户取消收藏文章
// @Description 用户取消收藏文章
// @Param data body ESIDListRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles/collects [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticlesApi) ArticleCollBatchRemoveView(c *gin.Context) {
	var cr ESIDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var articleIDList []string
	var collects []models.UserCollectModel

	global.DB.Find(&collects, "user_id = ? and article_id in ?", claims.UserID, cr.IDList).
		Select("article_id").Scan(&articleIDList)

	if len(articleIDList) == 0 {
		res.FailWithMessage("请求非法", c)
		return
	}

	var idList []interface{}
	for _, s := range articleIDList {
		idList = append(idList, s)
	}

	// global.Logger.Debugf("%T", idList...)  string
	// global.Logger.Debugf("%T", idList)     []interface{}
	boolSearch := elastic.NewTermsQuery("_id", idList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		count := article.CollectsCount - 1
		err = esServer.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": count,
		})
		if err != nil {
			global.Logger.Error(err)
			continue
		}
	}
	global.DB.Delete(&collects)
	res.OkWithMessage(fmt.Sprintf("成功取消收藏 %d 篇文章", len(articleIDList)), c)
}
