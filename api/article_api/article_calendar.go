package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"
)

type CalendarResp struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

// BucketType {"buckets":[{"key_as_string":"2024-12-18 00:00:00","key":1734480000000,"doc_count":4}]}
type BucketType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

// dataCount key time value cnt
var dateCount = map[string]int{}

// ArticleCalendarView 文章日历
// @Tags 文章管理
// @Summary 文章日历
// @Description 文章日历
// @Router /api/articles/calendar [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]CalendarResponse}
func (ArticlesApi) ArticleCalendarView(c *gin.Context) {
	// 时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")
	// 时间段搜索 1 year
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)
	// aYearAgo := now.Add(-100*time.Second) 一年前
	query := elastic.NewRangeQuery("created_at").
		Gte(aYearAgo.Format("2006-01-02 15:04:05")).
		Lte(now.Format("2006-01-02 15:04:05"))

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(100).
		Do(context.Background())

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	/*	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Logger.Error("解析文章数据出错:", err)
			continue
		}
		fmt.Printf("文章信息: %+v\n", article)
	}*/

	var data BucketType
	_ = json.Unmarshal(result.Aggregations["calendar"], &data)
	var resList = make([]CalendarResp, 0)
	for _, bucket := range data.Buckets {
		t, _ := time.Parse("2006-01-02 15:04:05", bucket.KeyAsString)
		// map key:time value:cnt
		dateCount[t.Format("2006-01-02")] = bucket.DocCount
		// 365天
		days := int(now.Sub(aYearAgo).Hours() / 24)
		for i := 0; i <= days; i++ {
			// 上一年第几天
			day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")
			// 发表数量
			count, _ := dateCount[day]
			resList = append(resList, CalendarResp{
				Date:  day,
				Count: count,
			})
		}
	}
	res.OkWithData(resList, c)

}
