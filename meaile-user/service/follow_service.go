package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
)

type FollowService interface {
	FollowUser(ctx *gin.Context, followBo bo.MeaileUserFollowBo) *model.Response
}
