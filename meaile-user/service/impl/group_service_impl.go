package impl

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/middlewares"
	_ "meaile-server/meaile-user/middlewares"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	"time"
)

type GroupServiceImpl struct {
}

func (u *GroupServiceImpl) SaveGroup(ctx *gin.Context, groupBo bo.MeaileFriendGroupBo) *model.Response {
	var group model.MeaileFriendGroup
	result := global.DB.Where(&model.MeaileFriendGroup{
		GroupName: groupBo.GroupName,
		UserId:    groupBo.UserId,
	}).First(&group)
	if result.RowsAffected == 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "分组信息已存在",
			Data: nil,
		}
	}
	token := ctx.Request.Header.Get("x-token")
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err,
		}
	}
	group = model.MeaileFriendGroup{
		GroupName:   groupBo.GroupName,
		UserId:      groupBo.UserId,
		CreatedBy:   customClaims.UserName,
		CreatedTime: time.Now(),
		UpdatedBy:   customClaims.UserName,
		UpdatedTime: time.Now(),
	}
	result = global.DB.Create(group)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "保存分组信息失败",
			Data: nil,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "保存分组信息成功",
		Data: group,
	}
}
func (u *GroupServiceImpl) DeleteGroup(ctx *gin.Context, groupIds bo.DeleteIds) *model.Response {
	if groupIds.GroupIds == nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "参数错误",
			Data: nil,
		}
	}
	result := global.DB.Where("id in (?)", groupIds.GroupIds).Delete(&model.MeaileFriendGroup{})
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "删除失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "删除成功",
		Data: nil,
	}
}
