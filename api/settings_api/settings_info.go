package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

type SettingUri struct {
	Name string `uri:"name"`
}

func (settingsApi *SettingsApi) SettingsInfoView(c *gin.Context) {

	var cr SettingUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMessage("没有对应配置信息", c)
	}
}
