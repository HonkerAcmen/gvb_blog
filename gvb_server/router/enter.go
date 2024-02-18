package router

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
	swaggerfile "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	router.GET("/swagger/*any", gs.WrapHandler(swaggerfile.Handler))
	apiRouterGroup := router.Group("api")
	// PublicGroup := Router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingRouter()
	return router
}
