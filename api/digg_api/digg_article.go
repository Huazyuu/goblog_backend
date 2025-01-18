package digg_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/service/redisServer"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" binding:"required" uri:"id"`
}

func (DiggApi) DiggArticleView(c *gin.Context) {
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
