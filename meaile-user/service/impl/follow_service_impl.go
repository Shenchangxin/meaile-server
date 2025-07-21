package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	result := global.DB.Where("user_name = ? and follow_user_name = ?", customClaims.UserName, followBo.FollowUserName).First(&followBo)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.SUCCESS,
			Msg:  "关注成功",
			Data: nil,
		}
	}
	follow := model.MeaileUserFollow{
		UserName:       customClaims.UserName,
		FollowUserName: followBo.FollowUserName,
		FollowTime:     time.Now(),
		CreatedBy:      customClaims.UserName,
		CreatedTime:    time.Now(),
		UpdatedBy:      customClaims.UserName,
		UpdatedTime:    time.Now(),
	}
	result = global.DB.Create(&follow)
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
}

func (f *FollowServiceImpl) UnfollowUser(ctx *gin.Context, followBo bo.MeaileUserFollowBo) *model.Response {
	claims, _ := ctx.Get("claims")
	customClaims := claims.(*model.CustomClaims)
	result := global.DB.Where("user_name = ? and follow_user_name = ?", customClaims.UserName, followBo.FollowUserName).First(&followBo)
	if result.RowsAffected == 1 {
		result = global.DB.Where("id = ?", followBo.Id).Delete(&followBo)
		if result.Error != nil {
			return &model.Response{
				Code: model.SUCCESS,
				Msg:  "取关成功",
				Data: followBo,
			}
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "取关成功",
		Data: nil,
	}
}
