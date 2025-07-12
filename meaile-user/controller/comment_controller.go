package controller

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
	"strconv"
)

func GetComments(ctx *gin.Context) {
	commentQuery := model.CommentQuery{}
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
	bizId := ctx.Query("bizId")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	commentQuery.PageNum = pageNum
	commentQuery.PageSize = pageSize
	commentQuery.BizId = bizId
	commentService := impl.CommentServiceImpl{}
	response := commentService.GetCommentList(ctx, commentQuery)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
