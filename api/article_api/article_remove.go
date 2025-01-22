package article_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/esServer"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

// ArticleRemoveView 删除文章
// @Tags 文章管理
// @Summary 删除文章
// @Description 删除文章
// @Param data body IDListRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticlesApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	bulkService := global.ESClient.Bulk().
		Index(models.ArticleModel{}.Index()).
		Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
		// full text
		go esServer.DeleteFullTextByArticleID(id)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("删除失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
}
