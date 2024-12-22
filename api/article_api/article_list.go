package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/esServer"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
}

func (ArticlesApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := esServer.CommList(esServer.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "abstract", "content"},
		Tag:      cr.Tag,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		res.FailWithMessage("查询失败", c)
		return
	}

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	// list 为空
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}
	// res.OkWithList(filter.Omit("list", list), int64(count), c)
	res.OkWithList(data, int64(count), c)
}
