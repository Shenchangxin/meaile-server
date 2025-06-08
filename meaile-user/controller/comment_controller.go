package controller

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
)

func GetComments(ctx *gin.Context) {
	commentQuery := model.CommentQuery{}
	if err := ctx.ShouldBind(&commentQuery); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	commentService := impl.CommentServiceImpl{}
	response := commentService.GetCommentList(ctx, commentQuery)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
