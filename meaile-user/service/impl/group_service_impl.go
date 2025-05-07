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
	token := ctx.Request.Header.Get("X-Token")
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
func (u *GroupServiceImpl) DeleteGroup(ctx *gin.Context, groupIds bo.DeleteGroupIds) *model.Response {
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
func (u *GroupServiceImpl) UpdateGroup(ctx *gin.Context, groupBo bo.MeaileFriendGroupBo) *model.Response {
	var group model.MeaileFriendGroup
	token := ctx.Request.Header.Get("X-Token")
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err,
		}
	}
	result := global.DB.Where(&model.MeaileFriendGroup{
		GroupName: groupBo.GroupName,
		UserId:    int64(customClaims.ID),
	}).First(&group)
	if result.RowsAffected == 1 {

		updateGroup := model.MeaileFriendGroup{
			GroupName:   groupBo.GroupName,
			UserId:      int64(customClaims.ID),
			UpdatedBy:   customClaims.UserName,
			UpdatedTime: time.Now(),
		}
		result := global.DB.Model(&model.MeaileFriendGroup{}).Where("id = ?", group.Id).Updates(&updateGroup)
		if result.Error != nil {
			return &model.Response{
				Code: model.FAILED,
				Msg:  "修改分组信息失败",
				Data: nil,
			}
		}
		return &model.Response{
			Code: model.SUCCESS,
			Msg:  "修改分组信息成功",
			Data: group,
		}
	} else {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "分组信息不存在",
			Data: nil,
		}
	}
}
func (u *GroupServiceImpl) GetGroupListByUserId(ctx *gin.Context, userId int64) *model.Response {
	var groups []model.MeaileFriendGroup
	result := global.DB.Where("user_id = ?", userId).Find(&groups)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询分组列表失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: groups,
	}
}
func (u *GroupServiceImpl) GetGroupById(ctx *gin.Context, groupId int64) *model.Response {
	var group model.MeaileFriendGroup
	result := global.DB.Where("id = ?", groupId).Find(&group)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询分组列表失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: group,
	}
}
