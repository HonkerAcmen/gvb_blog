package filesapi

type FilesAPI struct {
}

const (
	ImageFile FileType = "image"
	VideoFile FileType = "video"
	DocFile   FileType = "docfile"
)

type UploadFileInfo struct {
	FileName string   `json:"FileName"`
	IsUpload bool     `json:"isUpload"`
	Msg      string   `json:"msg"`
	Type     FileType `json:"type"`
}
