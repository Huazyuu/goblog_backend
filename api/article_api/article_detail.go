package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/service/esServer"
	"gvb_server/service/redisServer"
)

// ArticleDetailView 文章详情
// @Tags 文章管理
// @Summary 文章详情
// @Description 文章详情
// @Param id path string  true  "id"
// @Router /api/articles/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=models.ArticleModel}
func (ArticlesApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 增加look cnt
	_ = redisServer.NewArticleLook().Set(cr.ID)

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
