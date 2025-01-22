package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/service/redisServer"
)

// ArticleDiggView 文章点赞
// @Tags 文章管理
// @Summary 文章点赞
// @Description 文章点赞
// @Param data body models.ESIDRequest   true  "表示多个参数"
// @Router /api/articles/digg [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticlesApi) ArticleDiggView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 对长度校验
	// 查es
	err = redisServer.NewDigg().Set(cr.ID)
	if err != nil {
		res.FailWithError(err, &cr, c)
	}
	res.OkWithMessage("文章点赞成功", c)
}
