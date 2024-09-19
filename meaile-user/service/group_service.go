package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
)

type GroupService interface {
	SaveGroup(ctx *gin.Context, groupBo bo.MeaileFriendGroupBo) *model.Response
	UpdateGroup(ctx *gin.Context, groupBo bo.MeaileFriendGroupBo) *model.Response
	GetGroupListByUserId(ctx *gin.Context, userId int64) *model.Response
	GetGroupById(ctx *gin.Context, groupId int64) *model.Response
	DeleteGroup(ctx *gin.Context, groupIds bo.DeleteIds) *model.Response
}
