package app

import (
	"go-gin/app/common/request"
	"go-gin/app/common/response"
	"go-gin/app/services"
	"go-gin/cons"
	"go-gin/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *gin.Context) {
	// 请求参数校验
	var form request.Login
	if err := c.ShouldBind(&form); err != nil {
		errMsg := request.GetValidErrMsg(form, err)
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_BUSINESS_INVALID_PARAMS, errMsg, err, map[string]interface{}{
			cons.API_ERROR: err.Error(),
		})
		return
	}

	// 登录业务逻辑
	user, err := services.UserService.Login(form)
	if err != nil {
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_SERVER_LOGIN_FAILED, err.Error(), err, map[string]interface{}{
			cons.API_ERROR: err.Error(),
		})
		return
	} else {
		tokenData, _, err := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_SERVER_USER_FAILED_TOKEN, err.Error(), err, map[string]interface{}{
				cons.API_ERROR: err.Error(),
			})
			return
		}
		response.Success(c, tokenData)
	}
}

func GetUserInfo(c *gin.Context) {
	user, err := services.UserService.GetUserInfo(c.Keys[cons.API_USER_ID].(string))
	if err != nil {
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_SERVER_USER_OPERATION_FAILED, err.Error(), err, map[string]interface{}{
			cons.API_ERROR: err.Error(),
		})
		return
	}
	response.Success(c, user)
}

func Logout(c *gin.Context) {
	// 将 token 加入黑名单, 使其失效
	err := services.JwtService.JoinBlackList(c.Keys[cons.API_TOKEN_NAME].(*jwt.Token))
	if err != nil {
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_SERVER_USER_OPERATION_FAILED, err.Error(), err, map[string]interface{}{
			cons.API_ERROR: err.Error(),
		})
		return
	}
	response.Success(c, nil)
}
