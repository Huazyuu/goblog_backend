package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/esServer"
	"gvb_server/utils/jwt"
)

// ArticleCollectionCreateView 用户收藏文章，或取消收藏
// @Tags 文章管理
// @Summary 用户收藏文章，或取消收藏
// @Description 用户收藏文章，或取消收藏
// @Param data body ESIDRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles/collects [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticlesApi) ArticleCollectionCreateView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	model, err := esServer.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	var collection models.UserCollectModel
	err = global.DB.Take(&collection, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	var num = -1
	if err != nil {
		// 没有找到 收藏文章
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
		// 给文章的收藏数 +1
		num = 1
	}
	// 取消收藏
	global.DB.Delete(&collection)

	// 更新收藏
	err = esServer.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if num == 1 {
		res.OkWithMessage("收藏文章成功", c)
	} else {
		res.OkWithMessage("取消收藏成功", c)
	}
}
