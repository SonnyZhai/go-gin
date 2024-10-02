package request

type Register struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Email    string `form:"email" json:"email" binding:"required,email"`
}

// 自定义错误信息
func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "用户名不能为空",
		"password.required": "密码不能为空",
		"mobile.required":   "手机号不能为空",
		"mobile.mobile":     "手机号格式不正确",
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式不正确",
	}
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required,customValidatorUsername"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required":                "用户名不能为空",
		"username.customValidatorUsername": "用户名必须是有效的手机号或邮箱或用户名",
		"password.required":                "密码不能为空",
	}
}
