package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success = 0
	Err     = -1
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, c)
}

// 失败
func Fail(data any, msg string, c *gin.Context) {
	Result(Err, data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Err, map[string]interface{}{}, msg, c)
}

// 根据状态码返回信息
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	Result(Err, map[string]interface{}{}, "未知错误", c)
}
