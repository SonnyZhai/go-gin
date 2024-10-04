package bootstrap

import (
	"context"
	"go-gin/cons"
	"go-gin/global"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.uber.org/zap"
)

func InitializeS3() *s3.Client {
	// 连接 S3 服务
	r2, err := ConnectToS3(
		global.App.Config.Etcd.DefaultEndpoint,
		global.App.Config.Etcd.AccessKeyId,
		global.App.Config.Etcd.AccessKeySecret,
	)

	if err != nil {
		global.App.Log.Error(cons.ERROR_S3_CONNECTION, zap.Any("err", err))
		return nil
	}

	log.Printf(cons.INFO_S3_CONNECTION)

	// 检查并创建存储桶
	err = checkAndCreateBucket(r2, cons.OSS_R2_BUCKET_NAME, cons.OSS_R2_REGION)
	if err != nil {
		global.App.Log.Error("Failed to check or create bucket", zap.Any("err", err))
	}

	return r2
}

// ConnectToS3 连接 S3 服务
func ConnectToS3(endpoint, accessKey, secretAccessKey string) (*s3.Client, error) {
	// 连接 S3 服务
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secretAccessKey,
			"",
		)),
		config.WithRegion(cons.OSS_R2_REGION),
	)

	if err != nil {
		log.Fatalf(cons.FATAL_LOAD_R2_CONFIG+cons.STRING_PLACEHOLDER, err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	return client, nil
}

// 检查桶是否存在，不存在创建
func checkAndCreateBucket(client *s3.Client, bucketName, region string) error {
	// 检查存储桶是否存在
	_, err := client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		// 如果存储桶不存在，则创建存储桶
		_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraint(region),
			},
		})

		if err != nil {
			log.Printf("Couldn't create bucket %v in Region %v. Here's why: %v\n",
				bucketName, region, err)
			return err
		}

		log.Printf("Bucket %v created successfully in Region %v.\n", bucketName, region)
	} else {
		log.Printf("Bucket %v already exists.\n", bucketName)
	}

	return nil
}
