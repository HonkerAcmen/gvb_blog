package settingsapi

import (
	"gvb_server/common/res"
	"gvb_server/core"
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingsInfoUpdateView(c *gin.Context) {
	var ch = global.Config.SiteInfo
	err := c.ShouldBindJSON(&ch)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	global.Config.SiteInfo = ch
	err = core.UpdateConfigForSite()
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithMessage("成功", c)
}
