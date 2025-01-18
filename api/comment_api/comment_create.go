package comment_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redisServer"
	"gvb_server/utils/jwt"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论id
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	// 文章是否存在
	article := models.ArticleModel{}
	err = article.GetDataByID(cr.ArticleID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	if cr.ParentCommentID != nil {
		// 此评论为子评论
		// 父评论+1
		var parentComment models.CommentModel
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			res.FailWithMessage("父评论不存在", c)
			return
		}
		// 判断父评论的文章是否和当前文章一致
		if parentComment.ArticleID != cr.ArticleID {
			res.FailWithMessage("评论文章不一致", c)
			return
		}

		// +1
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}
	// update
	global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	})

	_ = redisServer.NewCommentCount().Set(cr.ArticleID)
	res.OkWithMessage("文章评论成功", c)
	return
}
