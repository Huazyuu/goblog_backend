package cornServer

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redisServer"
)

// SyncCommentData 同步评论数据到数据库
func SyncCommentData() {
	commentDiggInfo := redisServer.NewCommentDigg().GetInfo()
	for key, count := range commentDiggInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, key).Error
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		err = global.DB.Model(&comment).
			Update("digg_count", gorm.Expr("digg_count + ?", count)).Error
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		global.Logger.Infof("%s 更新成功 新的点赞数为：%d", comment.Content, comment.DiggCount)
	}
	redisServer.NewCommentDigg().Clear()
}
