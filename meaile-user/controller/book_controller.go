package controller

import (
	"github.com/gin-gonic/gin"
	model "meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
	"strconv"
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
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func UpdateBook(ctx *gin.Context) {
	bookBo := model.MeaileBookBo{}
	if err := ctx.ShouldBind(&bookBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	bookService := impl.BookServiceImpl{}
	response := bookService.UpdateBook(ctx, bookBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func GetMyBooks(ctx *gin.Context) {
	bookService := impl.BookServiceImpl{}
	response := bookService.GetMyBooks(ctx)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func GetBookListByTag(ctx *gin.Context) {
	bookBo := model.BookQueryBo{}
	tagIdStr := ctx.Query("tagId")
	sortField := ctx.Query("sortField")
	ascOrDesc := ctx.Query("ascOrDesc")
	tagId, err := strconv.ParseInt(tagIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	bookBo.TagId = tagId
	bookBo.SortField = sortField
	bookBo.AscOrDesc = ascOrDesc
	bookService := impl.BookServiceImpl{}
	response := bookService.GetBookListByTagId(ctx, bookBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func GetRecommendBookList(ctx *gin.Context) {
	bookBo := model.BookQueryBo{}
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
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
	bookBo.PageNum = pageNum
	bookBo.PageSize = pageSize

	bookService := impl.BookServiceImpl{}
	response := bookService.GetRecommendBookList(ctx, bookBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func DeleteBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	bookService := impl.BookServiceImpl{}
	response := bookService.DeleteBook(ctx, id)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func GetBookInfo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	bookService := impl.BookServiceImpl{}
	response := bookService.GetBookInfo(ctx, id)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
