package global

import (
	"fmt"
	"go-gin/cons"
)

// CustomError 定义自定义错误结构
type CustomError struct {
	HTTPStatus   int                    `json:"http_status"` // HTTP 状态码
	InternalCode cons.ErrorCode         `json:"code"`        // 错误码
	Message      string                 `json:"message"`     // 错误信息
	Context      map[string]interface{} `json:"context"`     // 错误上下文
}

// Error 实现 error 接口
func (e *CustomError) Error() string {
	return fmt.Sprintf("发生了错误： %d: %s", e.InternalCode, e.Message)
}

// WithContext 添加上下文信息
func (e *CustomError) WithContext(key string, value interface{}) *CustomError {
	e.Context[key] = value
	return e
}

// NewCustomError 创建新的自定义错误
func NewCustomError(code cons.ErrorCode, message string, httpStatus int, err error) *CustomError {
	return &CustomError{
		HTTPStatus:   httpStatus,
		InternalCode: code,
		Message:      message,
		Context:      make(map[string]interface{}),
	}
}
