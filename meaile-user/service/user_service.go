package service

import (
	"github.com/gin-gonic/gin"
	"meaile-web/meaile-user/model"
	bo "meaile-web/meaile-user/model/bo"
	vo "meaile-web/meaile-user/model/vo"
)

type UserService interface {
	GetUserInfo(ctx *gin.Context, userBo bo.MeaileUserBo) vo.MeaileUserVo
	GetUserList(ctx *gin.Context, userBo bo.MeaileUserBo) vo.MeaileUserVoList
	SaveUser(ctx *gin.Context, userBo bo.MeaileUserBo) (bool, error)
	UpdateUser(ctx *gin.Context, userBo bo.MeaileUserBo) bool
	Login(ctx *gin.Context, userBo bo.LoginForm) *model.Response
}
