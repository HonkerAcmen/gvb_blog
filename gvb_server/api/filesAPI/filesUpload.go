package filesapi

import (
	"errors"
	"gvb_server/common/res"
	"gvb_server/global"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	FileTypeMap = map[string]string{}
)

type UploadFileInfo struct {
	FileName string `json:"FileName"`
	IsUpload bool   `json:"isUpload"`
	Msg      string `json:"msg"`
	Type     string `json:"type"`
}

func GetFileType(fileName string) string {
	strSplits := strings.Split(fileName, ".")
	singerStr := strSplits[len(strSplits)-1:]
	str := strings.Join(singerStr, "")
	switch str {
	case "png", "jpg", "jpeg", "JPEG", "gif", "GIF", "bmp":
		return "image"
	case "mp4", "MKV", "mov", "wmv":
		return "video"
	case "doc", "pdf", "ppt", "xls", "xlsx", "docx", "txt":
		return "file"
	default:
		global.Log.Warn("未知类型文件")
		return ""
	}
}

// FileSizeIsOk: 判断文件的大小是否合理，合理返回true
func FileSizeIsOk(fileType string, size int64) (bool, error) {
	sizeAfter := float64(size) / float64(1024*1024)
	switch fileType {
	case "image":
		{
			return sizeAfter <= float64(global.Config.UploadInfo.PhotoSize), nil
		}
	case "video":
		{
			return sizeAfter <= float64(global.Config.UploadInfo.VideoSize), nil
		}
	case "file":
		{
			return sizeAfter <= float64(global.Config.UploadInfo.FileSize), nil
		}
	default:
		global.Log.Warn("未知文件, 获取大小失败")
		return false, errors.New("获取文件大小失败")
	}
}

// FilesUploadView: 上传多个文件
func (FilesAPI) FilesUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	FileList, ok := form.File["image"]

	if !ok {
		if err != nil {
			res.FailWithMessage("无效的图片", c)
			return
		}
	}
	for _, file := range FileList {
		var FilePath string
		fileType := GetFileType(file.Filename)
		switch fileType {
		case "image":
			{
				FilePath = path.Join("uploads/photos/" + file.Filename)
			}
		case "video":
			{
				FilePath = path.Join("uploads/videos/" + file.Filename)
			}
		case "file":
			{
				FilePath = path.Join("uploads/files/" + file.Filename)
			}
		default:
			FilePath = path.Join("uploads/unknown/" + file.Filename)
		}

		var resList []UploadFileInfo
		if isOK, err := FileSizeIsOk(fileType, file.Size); err != nil && !isOK {
			res.FailWithMessage(err.Error(), c)
			continue
		}

		resList = append(resList, UploadFileInfo{
			FileName: FilePath,
			IsUpload: true,
			Msg:      "上传成功",
			Type:     fileType,
		})

		// err := c.SaveUploadedFile(file, FilePath)
		// if err != nil {
		// 	global.Log.Warn("文件上传失败: ", err.Error())
		// 	res.FailWithCode(res.FileUploadErr, c)
		// 	return
		// }
		res.Ok(resList, "上传成功", c)
	}
}
