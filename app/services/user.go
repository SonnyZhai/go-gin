package services

import (
	"errors"
	"go-gin/app/common/request"
	"go-gin/app/models"
	"go-gin/cons"
	"go-gin/global"
	"go-gin/utils"
	"strconv"
)

// UserService 空结构体：用于实现用户服务
type userService struct {
}

// UserService 用户服务接口
var UserService = new(userService)

// Register 用户注册
func (u *userService) Register(params request.Register) (user models.User, err error) {
	// 查询用户是否存在
	var existingUser models.User
	// 是否存在指定username的用户，并只选择用户的 id 字段。
	result := global.App.DB.Where("username = ?", params.Username).Select("id").First(&existingUser)
	if result.Error == nil {
		return user, errors.New(cons.ERROR_USERNAME_EXIST)
	}

	// 检查手机号是否已存在
	result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&existingUser)
	if result.Error == nil {
		return user, errors.New(cons.ERROR_MOBILE_EXIST)
	}

	// 检查邮箱是否已存在
	result = global.App.DB.Where("email = ?", params.Email).Select("id").First(&existingUser)
	if result.Error == nil {
		return user, errors.New(cons.ERROR_EMAIL_EXIST)
	}

	// 创建用户
	user = models.User{
		Username: params.Username,
		Password: utils.BcryptMake([]byte(params.Password)),
		Mobile:   params.Mobile,
		Email:    params.Email,
	}

	// 保存用户信息到数据库
	result = global.App.DB.Create(&user)
	if result.Error != nil {
		return user, errors.New(cons.ERROR_REGISTER_FAILED)
	}

	return user, nil

}

// Login 登录
func (u *userService) Login(params request.Login) (user models.User, err error) {
	// 查询用户是否存在
	err = global.App.DB.Where("username = ? OR mobile = ? OR email = ?", params.Username, params.Username, params.Username).First(&user).Error
	// 校验密码
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New(cons.ERROR_LOGIN_FAILED)
	}
	return
}

// GetUserInfo 获取用户信息
func (u *userService) GetUserInfoByToken(id string) (user models.User, err error) {
	intId, convErr := strconv.Atoi(id)
	if convErr != nil {
		return user, errors.New(cons.ERROR_INVALID_USER_ID)
	}
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New(cons.ERROR_USER_NOT_EXIST)
	}
	return
}
