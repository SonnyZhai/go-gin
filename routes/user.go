package routes

import (
	"go-gin/app/common/request"
	"go-gin/app/common/response"
	"go-gin/cons"
	"go-gin/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * 用户分组路由
 */
func SetUserGroupRoutes(router *gin.RouterGroup) {
	router.POST("/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBind(&form); err != nil {
			errMsg := request.GetValidErrMsg(form, err)
			customErr := global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_PARAMS, errMsg, http.StatusBadRequest, err).
				WithContext("error", err.Error())
			c.JSON(http.StatusBadRequest, response.Fail(customErr))
			return
		}

		c.JSON(http.StatusOK, response.Success(form))
	})
}
