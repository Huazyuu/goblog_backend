package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var taglist []models.TagModel
	count := global.DB.Find(&taglist, cr.IDList).RowsAffected
	if count < 1 {
		res.FailWithMessage("标签不存在", c)
		return
	}
	// tag下有文章
	global.DB.Delete(&taglist)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)
}
