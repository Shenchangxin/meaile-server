package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	"mime/multipart"
)

type OssService interface {
	Upload(ctx *gin.Context, file *multipart.FileHeader, fileHeader *multipart.FileHeader) *model.Response
}
