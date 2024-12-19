package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type TagsResp struct {
	Tag           string   `json:"tag"`             // 文章标签
	Count         int      `json:"count"`           // 文章标签包含文章数
	ArticleIDList []string `json:"article_id_list"` // 文章title列表
	CreatedAt     string   `json:"created_at"`      // 创建时间
}
type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`       // tag name
		DocCount int    `json:"doc_count"` // 包含该文章tag的文章数量
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"` // article title
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

func (ArticlesApi) ArticleTagListView(c *gin.Context) {

	// 分页处理
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)
	if cr.Limit == 0 {
		cr.Limit = 10
	}
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	// 总数
	result, err := global.ESClient.Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")).
		Size(0).Do(context.Background())
	// 去重
	tmp, _ := result.Aggregations.Cardinality("tags")
	count := int64(*tmp.Value)

	// 查询tag下文章
	// [{"tag": "xxx","article_count": 2,"article_lists": []}]
	agg := elastic.NewTermsAggregation().Field("tags")
	// 标题子聚合
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	// 分页
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))
	query := elastic.NewBoolQuery()
	result, err = global.ESClient.Search(models.ArticleModel{}.Index()).
		Query(query).Aggregation("tags", agg).Size(0).Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	var tagType TagsType
	var tagList = make([]*TagsResp, 0)
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)

	var tagStringList = make([]string, 0) // 同步mysql tag创建的时间

	// es存储
	for _, bucket := range tagType.Buckets {
		var articleIDList []string
		for _, v := range bucket.Articles.Buckets {
			articleIDList = append(articleIDList, v.Key)
		}
		tagList = append(tagList, &TagsResp{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleIDList,
		})
		tagStringList = append(tagStringList, bucket.Key)
	}

	// 同步 mysql
	var tagModelList []models.TagModel
	global.DB.Find(&tagModelList, "title in ?", tagStringList)
	var tagDate = map[string]string{}
	for _, model := range tagModelList {
		tagDate[model.Title] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}
	for _, response := range tagList {
		response.CreatedAt = tagDate[response.Tag]
	}

	res.OkWithList(tagList, count, c)
}
