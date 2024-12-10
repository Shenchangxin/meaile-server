package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
)

type TagService interface {
	GetTagListByParentId(ctx *gin.Context, parentId int64) *model.Response
	GetTagListByUser(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response
	SaveTag(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response
	UpdateTag(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response
	DeleteTag(ctx *gin.Context, id int64) *model.Response
	//GetTagInfo(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response
}
