package services

import (
	"context"
	"go-gin/app/models"
	"go-gin/cons"
	"go-gin/global"
	"math/rand"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"
)

type fileService struct {
	presignClient *s3.PresignClient
}

var FileService = new(fileService)

func newFileService() *fileService {
	return &fileService{
		presignClient: s3.NewPresignClient(global.App.S3),
	}
}

func (f *fileService) UploadImages(uploadFile *multipart.FileHeader, uid, ext string) (file models.File, err error) {
	// 为文件生成唯一的文件名
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid := ulid.MustNew(ms, entropy)

	// 生成文件名
	objectKey := ulid.String() + cons.HYPHEN + uid + ext

	// 打开文件
	fileContent, err := uploadFile.Open()
	if err != nil {
		return file, err
	}
	defer fileContent.Close()

	// 上传文件
	_, err = global.App.S3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(cons.OSS_R2_BUCKET_NAME),
		Key:    aws.String(uid + cons.SLASH + objectKey),
		Body:   fileContent,
	})
	if err != nil {
		return file, err
	}
	// 保存文件信息到数据库

	return file, nil
}

// 检测用户文件夹是否存在
func CheckUserFolderExit(uid string) (err error) {
	// 检测用户文件夹是否存在
	_, err = global.App.S3.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(cons.OSS_R2_BUCKET_NAME),
		Key:    aws.String(uid + cons.SLASH),
	})

	// 如果用户文件夹不存在则创建
	if err != nil {
		err = createUserFolder(uid)
		if err != nil {
			global.App.Log.Error(cons.ERROR_CREATE_USER_FOLDER, zap.Any("err", err))
			return err
		}
	}
	return nil
}

// 根据用户的 ID 为用户创建一个文件夹
func createUserFolder(uid string) (err error) {
	// 创建用户文件夹
	resp, err := global.App.S3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(cons.OSS_R2_BUCKET_NAME),
		Key:    aws.String(uid + cons.SLASH),
	})

	if err != nil {
		global.App.Log.Error(cons.ERROR_CREATE_USER_FOLDER, zap.Any("err", err))
		return err
	}

	// 如果创建成功则返回
	global.App.Log.Info(cons.INFO_CREATE_USER_FOLDER, zap.Any("resp", resp))
	return nil
}
