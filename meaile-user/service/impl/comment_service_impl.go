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
	"strings"
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
	var creators []string
	for _, comment := range comments {
		creators = append(creators, comment.CreatedBy)
	}
	creatorsStr := strings.Join(creators, ", ")
	var users []vo.MeaileUserVo
	result = global.DB.Where("user_name in (?)", creatorsStr).Find(&users)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	for _, comment := range comments {
		for _, user := range users {
			if user.UserName == comment.CreatedBy {
				comment.Creator = user
			}
		}
	}
	return &model.Response{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: comments,
	}
}
