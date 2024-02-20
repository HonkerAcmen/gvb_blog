package router

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// router.GET("/swagger/*any", gs.WrapHandler(swaggerfile.Handler))

	// 路由分组
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}

	// 路由分层
	routerGroupApp.SettingRouter()
	routerGroupApp.DocFileRouter()
	routerGroupApp.ImageFileRouter()
	routerGroupApp.VideoFileRouter()
	return router
}
