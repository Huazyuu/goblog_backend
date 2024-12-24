package flag

import "gvb_server/models"

func EsCreateIndex() {
	// 文章
	_ = models.ArticleModel{}.CreateIndex()
	// 全文搜索
	_ = models.FullTextModel{}.CreateIndex()
}
