package docapi

import (
	filesapi "gvb_server/api/filesAPI"
	"gvb_server/common/res"
	"gvb_server/global"
	"path"

	"github.com/gin-gonic/gin"
)

func (DocFileAPI) DocFilesUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	FileList, ok := form.File[string(filesapi.DocFile)]
	if !ok {
		if err != nil {
			res.FailWithMessage("无效的文件", c)
			return
		}
	}
	for _, file := range FileList {
		filePath := path.Join("uploads/" + string(filesapi.DocFile) + "/" + file.Filename)
		var resLine []filesapi.UploadFileInfo
		fileType := filesapi.FilesAPI.GetFileType(filesapi.FilesAPI{}, filePath)
		if fileType != filesapi.DocFile {
			res.FailWithMessage("文件类型错误", c)
			continue
		}
		fileSize, err := filesapi.FilesAPI.GetFileSize(filesapi.FilesAPI{}, fileType, int64(global.Config.UploadInfo.FileSize))
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			continue
		}
		isOk, err := filesapi.FilesAPI.IsFileSize(filesapi.FilesAPI{}, fileSize, fileType)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			continue
		}
		if !isOk {
			res.FailWithMessage("文件超过了大小限制", c)
			continue
		}

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			res.FailWithCode(res.FileUploadErr, c)
			continue
		}

		resLine = append(resLine, filesapi.UploadFileInfo{
			FileName: filePath,
			IsUpload: true,
			Msg:      "上传成功",
			Type:     fileType,
		})
		res.OkWithData(resLine, c)
	}
}
