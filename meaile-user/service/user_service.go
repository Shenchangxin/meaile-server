package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	vo "meaile-server/meaile-user/model/vo"
)

type UserService interface {
	GetUserList(ctx *gin.Context, userBo bo.MeaileUserBo) vo.MeaileUserVoList
	SaveUser(ctx *gin.Context, userBo bo.MeaileUserBo) (bool, error)
	UpdateUser(ctx *gin.Context, userBo bo.MeaileUserBo) *model.Response
	GetUserFriendList(ctx *gin.Context) *model.Response
	Login(ctx *gin.Context, userBo bo.LoginForm) *model.Response
	GetUserInfo(ctx *gin.Context) *model.Response
	AddFriend(ctx *gin.Context, friendBo bo.AddUserFriendBo) *model.Response
	DeleteFriend(ctx *gin.Context, userId int64) *model.Response
}
