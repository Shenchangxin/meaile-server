package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/model"
	"mime/multipart"
	"strings"
)

type OssServiceImpl struct {
}

func (o *OssServiceImpl) Upload(ctx *gin.Context, file multipart.File, fileHeader *multipart.FileHeader) *model.Response {
	var u uuid.UUID
	var ossInfo model.MeaileOss
	var uuidStr string
	for {
		u = uuid.New()
		uuidStr = u.String()
		res := global.DB.Where(&model.MeaileOss{
			OssId: uuidStr,
		}).First(&ossInfo)
		if res.Error != nil {
			return &model.Response{
				Code: model.FAILED,
				Msg:  "文件上传失败",
				Data: res.Error,
			}
		}
		if res.RowsAffected == 0 {
			break
		}
	}
	err := global.MinioClient.UploadFile(global.ServerConfig.MinioConfig.BucketName, uuidStr, fileHeader)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "文件上传失败",
			Data: err,
		}
	}
	// 获取预签名URL（用于下载）
	//exp, err := time.ParseDuration("24h") // URL有效期为24小时
	//if err != nil {
	//	log.Fatalf("Failed to parse duration: %v", err)
	//}
	//url, err := global.MinioClient.GetPresignedGetObject(global.ServerConfig.MinioConfig.BucketName, uuidStr, exp)
	// 保存oss信息
	parts := strings.Split(fileHeader.Filename, ".")
	// 去掉切片中的最后一个元素
	if len(parts) > 1 {
		parts = parts[:len(parts)-1]
	} else {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "文件名称不正确",
			Data: nil,
		}
	}
	ossInfo = model.MeaileOss{
		OssId:    uuidStr,
		FileName: fileHeader.Filename,
		Suffix:   parts[0],
	}
	result := global.DB.Create(&ossInfo)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "保存文件信息失败",
			Data: nil,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "文件上传成功",
		Data: ossInfo,
	}
}
