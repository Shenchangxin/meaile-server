package impl

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	"time"
)

type FollowServiceImpl struct {
}

func (f *FollowServiceImpl) FollowUser(ctx *gin.Context, followBo bo.MeaileUserFollowBo) *model.Response {
	claims, _ := ctx.Get("claims")
	customClaims := claims.(*model.CustomClaims)
	follow := model.MeaileUserFollow{
		UserName:       customClaims.UserName,
		FollowUserName: followBo.FollowUserName,
		FollowTime:     time.Now(),
		CreatedBy:      customClaims.UserName,
		CreatedTime:    time.Now(),
		UpdatedBy:      customClaims.UserName,
		UpdatedTime:    time.Now(),
	}
	result := global.DB.Create(&follow)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "关注失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "关注成功",
		Data: follow,
	}

	return nil
}
