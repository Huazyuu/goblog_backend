package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/service/esServer"
	"gvb_server/service/redisServer"
)

func (ArticlesApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 增加look cnt
	redisServer.Look(cr.ID)

	model, err := esServer.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData(model, c)
}

type ESTitleRequest struct {
	Title string `json:"title" form:"title" binding:"required"`
}

func (ArticlesApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ESTitleRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	model, err := esServer.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
