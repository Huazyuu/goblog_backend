package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsInfoUpdateView 修改某一项的配置信
// SettingsInfoUpdateView 修改某一项的配置信息
// @Tags 系统管理
// @Summary 修改某一项的配置信息
// @Description 修改某一项的配置信息
// @Param name path int  true  "name"
// @Router /api/settings/{name} [put]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{}
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
