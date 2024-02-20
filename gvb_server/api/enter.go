package api

import (
	docapi "gvb_server/api/docAPI"
	imagesapi "gvb_server/api/imagesAPI"
	settingsapi "gvb_server/api/settings_api"
	videosapi "gvb_server/api/videosAPI"
)

type ApiGroup struct {
	SettingsApi  settingsapi.SettingApi
	DOCFileAPI   docapi.DocFileAPI
	ImageFileAPI imagesapi.ImageFileAPI
	VideoFileAPI videosapi.VideoFileAPI
}

var ApiGroupApp = new(ApiGroup)
