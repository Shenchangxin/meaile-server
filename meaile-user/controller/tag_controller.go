package controller

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/service/impl"
	"net/http"
	"strconv"
)

func TagListByParentId(ctx *gin.Context) {
	tagService := impl.TagServiceImpl{}
	idStr := ctx.Query("parentId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "参数错误",
			"data": err,
		})
		return
	}
	response := tagService.GetTagListByParentId(ctx, int64(id))
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
