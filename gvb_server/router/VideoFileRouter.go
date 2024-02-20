package router

import (
	"gvb_server/api"
	filesapi "gvb_server/api/filesAPI"
)

func (router RouterGroup) VideoFileRouter() {
	router.POST(string(filesapi.VideoFile), api.ApiGroupApp.VideoFileAPI.VideoFilesUpload)
}
