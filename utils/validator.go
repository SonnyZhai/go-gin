package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidateMobile 校验手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	return ok
}

// ValidateEmail 校验邮箱
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return ok
}

// CustomValidatorUsername 自定义用户名校验
// 用户名，手机号，邮箱三者必须满足其中一个
func CustomValidatorUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return ValidateMobile(fl) || ValidateEmail(fl) || len(username) > 0
}
