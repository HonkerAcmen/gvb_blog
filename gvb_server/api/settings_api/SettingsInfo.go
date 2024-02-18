package settingsapi

import (
	"gvb_server/common/res"

	"github.com/gin-gonic/gin"
)

func (Sapi SettingApi) SettingsInfoView(cxt *gin.Context) {
	res.FailWithCode(2, cxt)
}
