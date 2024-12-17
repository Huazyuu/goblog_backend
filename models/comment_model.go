package models

// CommentModel 评论表
type CommentModel struct {
	MODEL

	Content      string `gorm:"size:256" json:"content"`                // 评论内容
	DiggCount    int    `gorm:"size:8;default:0;" json:"digg_count"`    // 点赞数
	CommentCount int    `gorm:"size:8;default:0;" json:"comment_count"` // 子评论数

	UserID          uint   `json:"user_id"`                   // 评论的用户
	ArticleID       string `gorm:"size:32" json:"article_id"` // 文章i
	ParentCommentID *uint  `json:"parent_comment_id"`         // 父评论id

	// 外键
	// UserID -> UserModel(id)
	User UserModel `json:"user"` // 关联的用户
	// ArticleID -> ArticleModel(id)
	// Article ArticleModel `gorm:"foreignKey:ArticleID" json:"-"` // 关联的文章
	// ParentCommentID -> CommentModel(id)
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  // 子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` // 父级评论

}
