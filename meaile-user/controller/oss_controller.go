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
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func DownLoad(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数解析失败",
		})
		return
	}
	ossService := impl.OssServiceImpl{}

	response, fileName := ossService.Download(ctx, idInt)
	var data []byte

	// 使用类型断言来转换类型
	if byteData, ok := response.Data.([]byte); ok {
		data = byteData
		// 现在你可以使用 data 变量了，它已经是 []byte 类型
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "文件格式转换失败",
		})
		return
	}
	// 设置响应头
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", strconv.Itoa(len(data)))

	_, err = ctx.Writer.Write(data)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "文件写入失败",
		})
		return
	}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": response.Code,
	//	"msg":  response.Msg,
	//	"data": response.Data,
	//})
	//return
}

func GetUrl(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数解析失败",
		})
		return
	}
	ossService := impl.OssServiceImpl{}

	response := ossService.GetUrl(ctx, idInt)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
