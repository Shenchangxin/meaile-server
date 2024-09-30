package initialize

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/utils"
)

func InitMinio() {

	// 初始化 Minio 客户端
	minioClient, err := minio.New(global.ServerConfig.MinioConfig.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.ServerConfig.MinioConfig.AccessKey, global.ServerConfig.MinioConfig.SecretKey, ""),
		Secure: false, // 根据需要设置为 true 或 false
	})
	if err != nil {
		panic("初始化minio失败")
	}
	client := &utils.MinioClient{
		Client: minioClient,
	}
	global.MinioClient = client

}
