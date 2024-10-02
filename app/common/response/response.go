package response

import "go-gin/global"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// NewSuccessResponse 创建成功响应
func newSuccessResponse(data interface{}) Response {
	return Response{
		Code: 0,
		Msg:  "Success",
		Data: data,
	}
}

// NewFailResponse 创建失败响应
func newFailResponse(err *global.CustomError) Response {
	return Response{
		Code: 1,
		Msg:  "failed",
		Data: err,
	}
}

// Success 响应成功
func Success(data interface{}) Response {
	return newSuccessResponse(data)
}

// Fail 响应失败
func Fail(err *global.CustomError) Response {
	return newFailResponse(err)
}
