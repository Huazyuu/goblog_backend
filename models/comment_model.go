package models

// CommentModel 评论表
type CommentModel struct {
	// field
	MODEL           `json:",select(c)"`
	Content         string `gorm:"size:256" json:"content,select(c)"`                // 评论内容
	DiggCount       int    `gorm:"size:8;default:0;" json:"digg_count,select(c)"`    // 点赞数
	CommentCount    int    `gorm:"size:8;default:0;" json:"comment_count,select(c)"` // 子评论数
	UserID          uint   `json:"user_id,select(c)"`                                // 评论的用户
	ArticleID       string `gorm:"size:32" json:"article_id,select(c)"`              // 文章i
	ParentCommentID *uint  `json:"parent_comment_id,select(c)"`                      // 父评论id

	// gorm hand
	// 外键
	// UserID -> UserModel(id)
	User UserModel `json:"user,select(c)"` // 关联的用户
	// ArticleID -> ArticleModel(id)
	// Article ArticleModel `gorm:"foreignKey:ArticleID" json:"-"` // 关联的文章
	// ParentCommentID -> CommentModel(id)
	SubComments        []CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments,select(c)"` // 子评论列表
	ParentCommentModel *CommentModel  `gorm:"foreignKey:ParentCommentID" json:"comment_model"`          // 父级评论
}
