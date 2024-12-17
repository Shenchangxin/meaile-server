package utils

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"mime/multipart"
	"time"
)

type MinioClient struct {
	Client *minio.Client
}

// UploadFile 上传文件
func (m *MinioClient) UploadFile(bucketName, objectName string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("上传文件失败: %w", err)
	}
	// 上传文件到存储桶
	info, uploadErr := m.Client.PutObject(context.Background(), bucketName, objectName, src, file.Size, minio.PutObjectOptions{})
	if uploadErr != nil {
		return "", fmt.Errorf("上传文件失败: %w", err)
	}
	zap.S().Info("上传文件成功：", objectName)
	return info.Location, nil
}

// DownloadFile 下载文件
func (m *MinioClient) DownloadFile(bucketName, objectName, fileName string) (*minio.Object, error) {

	// 下载存储桶中的文件到本地
	obj, err := m.Client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	zap.S().Info("下载文件成功：", fileName)
	return obj, nil

}

// DeleteFile 删除文件
func (m *MinioClient) DeleteFile(bucketName, objectName string) (bool, error) {
	// 删除存储桶中的文件
	err := m.Client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		zap.S().Info("删除文件失败：", objectName)
		return false, err
	}

	zap.S().Info("删除文件成功：", objectName)
	return true, err
}

// ListObjects 列出文件
func (m *MinioClient) ListObjects(bucketName string) ([]string, error) {
	var objectNames []string

	for object := range m.Client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{}) {
		if object.Err != nil {
			return nil, object.Err
		}

		objectNames = append(objectNames, object.Key)
	}

	return objectNames, nil
}

// GetPresignedGetObject 返回对象的url地址，有效期时间为expires
func (m *MinioClient) GetPresignedGetObject(bucketName string, objectName string, expires time.Duration) (string, error) {
	object, err := m.Client.PresignedGetObject(context.Background(), bucketName, objectName, expires, nil)
	if err != nil {
		zap.S().Error("获取对象的url失败", err)
		return "", err
	}

	return object.String(), nil
}
