package data_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"
)

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
type DateCountResponse struct {
	DateList   []string `json:"date_list"`
	LoginCount []int    `json:"login_count"`
	SignCount  []int    `json:"sign_count"`
}

// SevenLoginView 七天登录用户统计
func (DataApi) SevenLoginView(c *gin.Context) {
	var loginDateCount, signDateCount []DateCount
	// k-v: data count
	var loginDateCountMap, signDateCountMap = make(map[string]int), make(map[string]int)
	var loginCountList, signCountList []int
	var dateList []string
	now := time.Now()

	// 查询并scan到list
	global.DB.Model(models.LoginDataModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&loginDateCount)
	global.DB.Model(models.UserModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&signDateCount)

	for _, date := range loginDateCount {
		loginDateCountMap[date.Date] = date.Count
	}
	for _, date := range signDateCount {
		signDateCountMap[date.Date] = date.Count
	}

	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		loginCount := loginDateCountMap[day]
		signCount := signDateCountMap[day]
		dateList = append(dateList, day)
		loginCountList = append(loginCountList, loginCount)
		signCountList = append(signCountList, signCount)
	}
	res.OkWithData(DateCountResponse{
		DateList:   dateList,
		LoginCount: loginCountList,
		SignCount:  signCountList,
	}, c)
}
