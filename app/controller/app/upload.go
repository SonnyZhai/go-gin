package app

import (
	"go-gin/app/common/request"
	"go-gin/app/middleware"
	"go-gin/app/services"
	"go-gin/cons"
	"go-gin/errors"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	// 请求参数校验
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		errMsg := request.GetValidErrMsg(form, err)
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_BUSINESS_INVALID_IMAGE,
			errMsg, err, map[string]interface{}{
				cons.API_ERROR: err.Error(),
			})
		return
	}

	// 获取上传文件
	formData, err := c.MultipartForm()
	if err != nil {
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_BUSINESS_INVALID_IMAGE,
			cons.ERROR_UPLOAD_IMAGE_RECEIVED, err, map[string]interface{}{
				cons.API_ERROR: err.Error(),
			})
		return
	}

	// 检查是否有文件
	files := formData.File[cons.FILE_TYPE]
	if len(files) == 0 {
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_BUSINESS_INVALID_IMAGE,
			cons.ERROR_UPLOAD_IMAGE_RECEIVED, nil, nil)
		return
	}

	for _, file := range files {
		// 验证文件大小
		if file.Size > 2048*2048 {
			errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_BUSINESS_INVALID_IMAGE,
				cons.ERROR_UPLOAD_IMAGE_SIZE, nil, nil)
			return
		}

		// 验证文件格式
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_BUSINESS_INVALID_IMAGE,
				cons.ERROR_UPLOAD_IMAGE_FORMAT, nil, nil)
			return
		}

		// 保存文件业务逻辑，保存到s3
		uid := CheckUserFolder(c)
		file, err := services.FileService.UploadImages(file, uid, ext)

	}

}

func CheckUserFolder(c *gin.Context) (uid string) {
	uid, ok := middleware.GetUserID(c)
	if !ok {
		errors.HandleErrorWithContext(c, http.StatusBadRequest, cons.ERROR_CODE_BUSINESS_INVALID_PARAMS,
			cons.ERROR_USER_ID_EMPTY, nil, nil)
		return
	}

	// 检测用户文件夹是否存在
	err := services.CheckUserFolderExit(uid)
	if err != nil {
		errors.HandleErrorWithContext(c, http.StatusInternalServerError, cons.ERROR_CODE_SERVER_UNKNOWN,
			cons.ERROR_UNKNOWN_SERVER_ERROR, err, nil)
		return
	}

	return uid
}
