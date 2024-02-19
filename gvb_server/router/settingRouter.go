package router

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingRouter() {
	settingApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings", settingApi.SettingsInfoView)
	router.PUT("/settings", settingApi.SettingsInfoUpdateView)
}
