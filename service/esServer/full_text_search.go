package esServer

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"gvb_server/global"
	"gvb_server/models"
	"strings"
)

type SearchData struct {
	Key   string `json:"key"`
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
	Body  string `json:"body"`  // 正文
}

func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")

	var isCodeBlock = false // 是否代码块
	var headList, bodyList []string
	var body string

	headList = append(headList, getHeader(title))

	for _, line := range dataList {
		if strings.HasPrefix(line, "```") {
			// 是代码块
			isCodeBlock = !isCodeBlock
		}
		// 如果行是标题（包含 # 且不在代码块中）：
		if strings.Contains(line, "#") && !isCodeBlock {
			// 下一篇标题
			headList = append(headList, getHeader(line))
			// 这篇文章内容
			bodyList = append(bodyList, getBody(body))
			body = "" // 清空 body，为下一个正文内容做准备。
			continue
		}
		body += line
	}
	// 循环结束后，添加最后一个正文内容到 bodyList。
	bodyList = append(bodyList, getBody(body))

	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
			Key:   id,
		})
	}
	return searchDataList
}

// getHeader 删除标题中的 # 和空格，返回处理后的标题字符串。
func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

// getBody 使用 blackfriday 将正文（Markdown 格式）解析为 HTML
// 使用 goquery 从 HTML 提取纯文本
func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

// getSlug 生成跳转地址
func getSlug(slug string) string {
	return "#" + slug
}

// ASyncArticleFullText 异步同步全文搜索
func ASyncArticleFullText(id, title, content string) {
	indexList := GetSearchIndexDataByContent(id, title, content)
	/*批量添加
	遍历 indexList，即需要添加到 Elasticsearch 的数据。
	elastic.NewBulkIndexRequest()每一条数据都会生成一个 BulkIndexRequest：
	.Index(models.FullTextModel{}.Index())：指定写入的索引（类似数据库中的表名）。
	.Doc(indexData)：指定要写入的数据内容（通常是 JSON 格式的文档）。
	bulk.Add(req)：将生成的 BulkIndexRequest 添加到 Bulk 请求中。*/
	bulk := global.ESClient.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	global.Logger.Infof("%s 添加成功,共 %d条", title, len(result.Succeeded()))
}

// DeleteFullTextByArticleID 删除全文搜索数据
func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := global.ESClient.
		DeleteByQuery().Index(models.FullTextModel{}.Index()).
		Query(boolSearch).Do(context.Background())
	global.Logger.Infof("成功删除 %d 条记录", res.Deleted)
}
