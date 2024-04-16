package api

import "gvb_server/api/settings_api"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}
