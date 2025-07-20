package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	vo "meaile-server/meaile-user/model/vo"
	"net/http"
)

type CommentServiceImpl struct {
}

func (c *CommentServiceImpl) GetCommentList(ctx *gin.Context, query bo.CommentQuery) *model.Response {
	var comments []vo.MeaileCommentVo
	offset := (query.PageNum - 1) * query.PageSize
	db := global.DB.Offset(offset).Limit(query.PageSize)
	if query.BizId != "" {
		db.Where("biz_id = ?", query.BizId)
	}
	result := db.Order("created_time DESC").Find(&comments)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	var creatorsIds []int64
	for _, comment := range comments {
		creatorsIds = append(creatorsIds, comment.UserId)
	}
	var users []vo.MeaileUserVo
	result = global.DB.Where("id in (?)", creatorsIds).Find(&users)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	var imageOssIds []string
	for _, user := range users {
		imageOssIds = append(imageOssIds, user.Avatar)
	}
	var ossList []model.MeaileOss
	result = global.DB.Where("oss_id in (?)", imageOssIds).Find(&ossList)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	for i, user := range users {
		for _, oss := range ossList {
			if oss.OssId == user.Avatar {
				fileUrl := global.ServerConfig.HuaWeiOBSConfig.UrlPrefix + oss.FileName
				users[i].AvatarUrl = fileUrl
				break
			}
		}
	}
	for i, comment := range comments {
		for _, user := range users {
			if user.Id == comment.UserId {
				comments[i].AvatarUrl = user.AvatarUrl
				comments[i].UserName = user.UserName
				comments[i].Creator = user
			}
		}
	}
	return &model.Response{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: comments,
	}
}
