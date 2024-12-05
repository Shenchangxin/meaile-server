package controller

import (
	"github.com/gin-gonic/gin"
	model "meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
)

func SaveBook(ctx *gin.Context) {
	bookBo := model.MeaileBookBo{}
	if err := ctx.ShouldBind(&bookBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	bookService := impl.BookServiceImpl{}
	response := bookService.SaveBook(ctx, bookBo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func GetBookListByTag(ctx *gin.Context) {
	bookBo := model.BookQueryBo{}
	if err := ctx.ShouldBind(&bookBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	bookService := impl.BookServiceImpl{}
	response := bookService.GetBookListByTagId(ctx, bookBo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
