package res

type ErrorCode int

const (
	SystemError    ErrorCode = 1001 // 系统错误
	ParameterError ErrorCode = 1002 // 参数错误
	FileUploadErr  ErrorCode = 1003 // 文件上传错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SystemError:    "系统错误",
		ParameterError: "参数错误",
		FileUploadErr:  "文件上传错误",
	}
)
