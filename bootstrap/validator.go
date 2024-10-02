package bootstrap

import (
	"go-gin/cons"
	"go-gin/utils"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitializeValidator() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		v.RegisterValidation("mobile", utils.ValidateMobile)
		v.RegisterValidation("email", utils.ValidateEmail)
		v.RegisterValidation("customValidatorUsername", utils.CustomValidatorUsername)

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get(cons.JSON_TYPE), cons.COMMA, 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
