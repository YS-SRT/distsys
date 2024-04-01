package utils

import (
	"context"
	"distsys/global"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

func UploadFile(objectName string, reader io.Reader, objectSize int64) (ok bool) {

	info, err := global.GMinioCli.PutObject(context.TODO(), global.GCONFIG.MinioAuth.Bucket, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		zap.L().Error("minio put data error: ", zap.Error(err))
		return false
	}
	zap.L().Sugar().Info("Successfully uploaded bytes: ", info)
	return true
}

func GetFileUrl(fileName string, expires time.Duration) string {

	reqParams := make(url.Values)
	presignedURL, err := global.GMinioCli.PresignedGetObject(context.TODO(), global.GCONFIG.MinioAuth.Bucket, fileName, expires, reqParams)
	if err != nil {
		zap.L().Error(err.Error())
		return ""
	}
	zap.L().Sugar().Info("got minio url paht: %s", presignedURL)
	return presignedURL.RawPath
}
