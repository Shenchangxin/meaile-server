package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	"mime/multipart"
)

type OssService interface {
	Upload(ctx *gin.Context, fileHeader *multipart.FileHeader) *model.Response
	Download(ctx *gin.Context, id int64) *model.Response
	GetUrl(ctx *gin.Context, id int64) *model.Response
}
