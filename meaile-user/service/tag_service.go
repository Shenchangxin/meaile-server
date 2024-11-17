package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
)

type TagService interface {
	GetTagListByParentId(ctx *gin.Context, parentId int64) *model.Response
	//GetTagListByUser(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response
	//GetTagInfo(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response
}
