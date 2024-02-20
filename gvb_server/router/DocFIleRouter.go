package router

import (
	"gvb_server/api"
	filesapi "gvb_server/api/filesAPI"
)

func (router RouterGroup) DocFileRouter() {
	router.POST(string(filesapi.DocFile), api.ApiGroupApp.DOCFileAPI.DocFilesUpload)
}
