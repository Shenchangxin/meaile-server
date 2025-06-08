package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
)

type CommentService interface {
	GetCommentList(ctx *gin.Context, query bo.CommentQuery) *model.Response
}
