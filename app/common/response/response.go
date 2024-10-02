package response

import (
	"go-gin/cons"
	"go-gin/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// NewSuccessResponse 创建成功响应
func newSuccessResponse(data interface{}) Response {
	return Response{
		Code: 0,
		Msg:  cons.API_SUCCESS,
		Data: data,
	}
}

// NewFailResponse 创建失败响应
func newFailResponse(err *global.CustomError) Response {
	return Response{
		Code: 1,
		Msg:  cons.API_FAILED,
		Data: err,
	}
}

// Success 响应成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, newSuccessResponse(data))
}

// Fail 响应失败
func Fail(c *gin.Context, err *global.CustomError) {
	c.JSON(err.HTTPStatus, newFailResponse(err))
}
