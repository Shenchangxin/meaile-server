package impl

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	"mime/multipart"
)

type OssServiceImpl struct {
}

func (o *OssServiceImpl) Upload(ctx *gin.Context, file *multipart.FileHeader, fileHeader *multipart.FileHeader) *model.Response {
	return nil
}
