package controller

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/service/impl"
	"net/http"
	"strconv"
)

func Upload(ctx *gin.Context) {

	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "文件上传失败",
		})
		return
	}
	ossService := impl.OssServiceImpl{}

	response := ossService.Upload(ctx, header)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func DownLoad(ctx *gin.Context) {
	id := ctx.Query("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数解析失败",
		})
		return
	}
	ossService := impl.OssServiceImpl{}

	response := ossService.Download(ctx, idInt)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func GetUrl(ctx *gin.Context) {
	id := ctx.Query("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数解析失败",
		})
		return
	}
	ossService := impl.OssServiceImpl{}

	response := ossService.GetUrl(ctx, idInt)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
