package router

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.FilesAPI
	router.POST("images", imagesApi.FilesUploadView)
}
