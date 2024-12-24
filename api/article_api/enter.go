package article_api

type ArticlesApi struct {
}
type ESIDRequest struct {
	ID string `json:"id" form:"id" binding:"required" uri:"id"`
}
