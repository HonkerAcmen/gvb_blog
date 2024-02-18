package api

import settingsapi "gvb_server/api/settings_api"

type ApiGroup struct {
	SettingsApi settingsapi.SettingApi
}

var ApiGroupApp = new(ApiGroup)
