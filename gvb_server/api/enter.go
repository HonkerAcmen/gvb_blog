package api

import (
	filesapi "gvb_server/api/filesAPI"
	settingsapi "gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settingsapi.SettingApi
	FilesAPI    filesapi.FilesAPI
}

var ApiGroupApp = new(ApiGroup)
