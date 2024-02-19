package settingsapi

import (
	"gvb_server/common/res"
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

func (Sapi SettingApi) SettingsInfoView(cxt *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, cxt)
}
