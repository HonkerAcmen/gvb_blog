package filesapi

import (
	"errors"
	"gvb_server/global"
	"strings"
)

type FileType string

// GetFileType: 获取文件的类型
func (FilesAPI) GetFileType(fileName string) FileType {
	strSplits := strings.Split(fileName, ".")
	singerStr := strSplits[len(strSplits)-1:]
	str := strings.Join(singerStr, "")
	switch str {
	case "png", "jpg", "jpeg", "JPEG", "gif", "GIF", "bmp":
		return ImageFile
	case "mp4", "MKV", "mov", "wmv":
		return VideoFile
	case "doc", "pdf", "ppt", "xls", "xlsx", "docx", "txt":
		return DocFile
	default:
		global.Log.Warn("未知类型文件")
		return ""
	}
}

// GetFileSize: 获取文件大小
func (FilesAPI) GetFileSize(fileType FileType, size int64) (float64, error) {
	sizeAfter := float64(size) / float64(1024*1024)
	switch fileType {
	case ImageFile:
		{
			return sizeAfter, nil
		}
	case VideoFile:
		{
			return sizeAfter, nil
		}
	case DocFile:
		{
			return sizeAfter, nil
		}
	default:
		global.Log.Warn("未知文件, 获取大小失败")
		return 0, errors.New("获取文件大小失败")
	}
}

// IsFileSize: 判断文件大小是否合理
func (FilesAPI) IsFileSize(fileSize float64, fileType FileType) (isOk bool, err error) {

	switch fileType {
	case ImageFile:
		{
			isOk = fileSize <= float64(global.Config.UploadInfo.PhotoSize)
		}
	case VideoFile:
		{
			isOk = fileSize <= float64(global.Config.UploadInfo.VideoSize)
		}
	case DocFile:
		{
			isOk = fileSize <= float64(global.Config.UploadInfo.FileSize)
		}
	default:
		global.Log.Warn("未知类型文件，无法判断大小")
		isOk = false
		err = errors.New("未知类型文件，无法判断大小")
		return isOk, err
	}
	return isOk, nil
}

// // FilesUploadView: 上传多个文件
// func (FilesAPI) FilesUploadView(c *gin.Context) {
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		res.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	FileList, ok := form.File["image"]

// 	if !ok {
// 		if err != nil {
// 			res.FailWithMessage("无效的图片", c)
// 			return
// 		}
// 	}
// 	for _, file := range FileList {
// 		var FilePath string
// 		fileType := GetFileType(file.Filename)
// 		switch fileType {
// 		case "image":
// 			{
// 				FilePath = path.Join("uploads/photos/" + file.Filename)
// 			}
// 		case "video":
// 			{
// 				FilePath = path.Join("uploads/videos/" + file.Filename)
// 			}
// 		case "file":
// 			{
// 				FilePath = path.Join("uploads/files/" + file.Filename)
// 			}
// 		default:
// 			FilePath = path.Join("uploads/unknown/" + file.Filename)
// 		}

// 		var resList []UploadFileInfo
// 		fileSize, err := GetFileSize(fileType, file.Size)
// 		if err != nil {
// 			res.FailWithMessage(err.Error(), c)
// 			continue
// 		}
// 		isOk, err := FileSizeIsOk(fileSize, fileType)
// 		if err != nil {
// 			res.FailWithMessage(err.Error(), c)
// 			continue
// 		}
// 		if !isOk {
// 			res.FailWithMessage("文件太大!", c)
// 			continue
// 		}

// 		err = c.SaveUploadedFile(file, FilePath)
// 		if err != nil {
// 			global.Log.Warn("文件上传失败: ", err.Error())
// 			res.FailWithCode(res.FileUploadErr, c)
// 			return
// 		}
// 		resList = append(resList, UploadFileInfo{
// 			FileName: FilePath,
// 			IsUpload: true,
// 			Msg:      "上传成功",
// 			Type:     fileType,
// 		})
// 		res.OkWithData(resList, c)
// 	}
// }
