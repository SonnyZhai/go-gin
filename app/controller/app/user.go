package app

import (
	"go-gin/app/common/request"
	"go-gin/app/common/response"
	"go-gin/app/services"
	"go-gin/cons"
	"go-gin/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// 请求参数校验
	var form request.Register
	if err := c.ShouldBind(&form); err != nil {
		errMsg := request.GetValidErrMsg(form, err)
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_BUSINESS_INVALID_PARAMS, errMsg, err, map[string]interface{}{
			cons.API_ERROR: err.Error(),
		})
		return
	}

	// 注册用户业务逻辑
	user, err := services.UserService.Register(form)
	if err != nil {
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_SERVER_REGISTER_FAILED, err.Error(), err, map[string]interface{}{
			cons.API_ERROR: err.Error(),
		})
		return
	}

	response.Success(c, user)
}
