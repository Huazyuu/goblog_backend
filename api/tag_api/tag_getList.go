package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	// todo 显示tag下的标签数
	res.OkWithList(list, count, c)
}
