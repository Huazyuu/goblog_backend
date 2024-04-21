package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (settingsApi *SettingsApi) SettingsInfoUpdateView(c *gin.Context) {

	var cr SettingUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	default:
		res.FailWithMessage("没有对应配置信息", c)
		return
	}

	if err = core.SetYaml(); err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithSuccess(c)
}
