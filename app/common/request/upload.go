package request

import "mime/multipart"

type ImageUpload struct {
	Image *multipart.FileHeader `form:"image" json:"image" binding:"required"`
}

func (imageUpload ImageUpload) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"image.required": "请选择要上传的图片",
	}
}
