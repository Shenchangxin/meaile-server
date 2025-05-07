package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/middlewares"
	"meaile-server/meaile-user/model"
	"mime/multipart"
	"strings"
	"time"
)

type OssServiceImpl struct {
}

func (o *OssServiceImpl) Upload(ctx *gin.Context, fileHeader *multipart.FileHeader) *model.Response {
	var u uuid.UUID
	var ossInfo model.MeaileOss
	var uuidStr string
	token := ctx.Request.Header.Get("X-Token")
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败",
			Data: err,
		}
	}
	parts := strings.Split(fileHeader.Filename, ".")
	if len(parts) <= 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "文件名称不正确",
			Data: nil,
		}
	}
	for {
		u = uuid.New()
		uuidStr = u.String()
		res := global.DB.Where(&model.MeaileOss{
			OssId: uuidStr,
		}).First(&ossInfo)
		if res.Error != nil {
			if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return &model.Response{
					Code: model.FAILED,
					Msg:  "文件上传失败",
					Data: res.Error,
				}
			}
		}
		if res.RowsAffected == 0 {
			break
		}
	}
	objectName := uuidStr + "." + parts[len(parts)-1]
	location, uploadErr := global.MinioClient.UploadFile(global.ServerConfig.MinioConfig.BucketName, objectName, fileHeader)
	if uploadErr != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "文件上传失败",
			Data: uploadErr,
		}
	}
	// 获取预签名URL（用于下载）
	//exp, err := time.ParseDuration("24h") // URL有效期为24小时
	//if err != nil {
	//	log.Fatalf("Failed to parse duration: %v", err)
	//}
	//url, err := global.MinioClient.GetPresignedGetObject(global.ServerConfig.MinioConfig.BucketName, uuidStr, exp)
	// 保存oss信息

	ossInfo = model.MeaileOss{
		OssId:       uuidStr,
		FileName:    fileHeader.Filename,
		Suffix:      "." + parts[len(parts)-1],
		FileUrl:     location,
		CreatedTime: time.Now(),
		CreatedBy:   customClaims.UserName,
		UpdatedTime: time.Now(),
		UpdatedBy:   customClaims.UserName,
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
func (o *OssServiceImpl) Download(ctx *gin.Context, id int64) (res *model.Response, fileName string) {
	var ossInfo model.MeaileOss
	result := global.DB.Where(&model.MeaileOss{
		Id: id,
	}).First(&ossInfo)

	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "下载失败",
			Data: result.Error,
		}, ""
	}
	if result.RowsAffected != 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "未找到该文件信息",
			Data: nil,
		}, ""
	}
	object, err := global.MinioClient.DownloadFile(global.ServerConfig.MinioConfig.BucketName, ossInfo.OssId+ossInfo.Suffix, ossInfo.FileName)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "下载失败",
			Data: result.Error,
		}, ""
	}
	defer object.Close()

	// 读取对象内容
	data, err := io.ReadAll(object)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "下载失败",
			Data: result.Error,
		}, ""
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "下载成功",
		Data: data,
	}, ossInfo.FileName
}
func (o *OssServiceImpl) GetUrl(ctx *gin.Context, id int64) *model.Response {
	var ossInfo model.MeaileOss
	result := global.DB.Where(&model.MeaileOss{
		Id: id,
	}).First(&ossInfo)

	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "下载失败",
			Data: result.Error,
		}
	}
	if result.RowsAffected != 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "未找到该文件信息",
			Data: nil,
		}
	}
	// 获取预签名URL（用于下载）
	exp, err := time.ParseDuration("24h") // URL有效期为24小时
	if err != nil {
		zap.S().Fatalf("获取时间失败: %v", err)
		return &model.Response{
			Code: model.FAILED,
			Msg:  "操作失败",
			Data: err,
		}
	}
	res, err := global.MinioClient.GetPresignedGetObject(global.ServerConfig.MinioConfig.BucketName, ossInfo.OssId+ossInfo.Suffix, exp)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "操作失败",
			Data: result.Error,
		}
	}

	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: res,
	}
}
