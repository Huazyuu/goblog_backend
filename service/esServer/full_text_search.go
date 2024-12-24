package esServer

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

type SearchData struct {
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
