package request

import (
	"go-gin/cons"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

/**
 * 从验证错误中提取错误信息，并返回一个字符串。支持自定义错误信息，并且会将所有错误信息连接成一个字符串返回。如果没有错误信息，则返回默认错误信息。
 * @param request interface 请求参数
 * @param err error 错误信息
 * @return string 错误信息
 */
func GetValidErrMsg(request interface{}, err error) string {
	// 参数校验错误
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, 0)            // 错误信息列表
		validator, isValidator := request.(Validator) // 是否实现了 Validator 接口

		// 遍历错误信息
		for _, v := range validationErrors {
			// 如果实现了 Validator 接口，则尝试获取自定义错误信息
			if isValidator {
				if message, exist := validator.GetMessages()[v.Field()+cons.DOT+v.Tag()]; exist {
					errorMessages = append(errorMessages, message)
					continue
				}
			}
			errorMessages = append(errorMessages, v.Error())
		}

		// 返回错误信息
		if len(errorMessages) > 0 {
			return strings.Join(errorMessages, cons.SEMICOLON)
		}
	}

	// 默认错误信息
	return cons.ERROR_DEFAULT_REQUEST
}
