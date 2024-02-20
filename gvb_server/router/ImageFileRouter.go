package router

import (
	"gvb_server/api"
	filesapi "gvb_server/api/filesAPI"
)

func (router RouterGroup) ImageFileRouter() {
	router.POST(string(filesapi.ImageFile), api.ApiGroupApp.ImageFileAPI.ImageFilesUpload)
}
