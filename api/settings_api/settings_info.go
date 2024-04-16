package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

func (SettingsApi *SettingsApi) SettingsInfoView(c *gin.Context) {
	res.Ok(
		map[string]string{
			"id":     "123",
			"name":   "root",
			"action": "test"},
		"res.ok", c)
	//res.OkWithData(map[string]string{"okWithData": "success"}, c)
	//res.OkWithMessage("res.okWithMessage", c)
	//res.Fail(map[string]string{}, "res.fail", c)
	//res.FailWithMessage("res.failWithMag", c)
	//res.FailWithCode(res.SettingsError, c)
}
