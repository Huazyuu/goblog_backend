package flag

import "gvb_server/models"

func EsCreateIndex() {
	// 文章
	// _ = models.ArticleModel{}.CreateIndex()
	// todo debug全文搜索 防止删除文章 先注释掉
	_ = models.FullTextModel{}.CreateIndex()
}
