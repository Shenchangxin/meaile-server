package controller

import (
	"github.com/gin-gonic/gin"
	bo "meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
	"strconv"
)

func TagListByParentId(ctx *gin.Context) {
	tagService := impl.TagServiceImpl{}
	idStr := ctx.Query("parentId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "参数错误",
			"data": err,
		})
		return
	}
	response := tagService.GetTagListByParentId(ctx, int64(id))
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func TagListByUser(ctx *gin.Context) {
	tagService := impl.TagServiceImpl{}
	idStr := ctx.Query("parentId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "参数错误",
			"data": err,
		})
		return
	}
	var tagBo bo.MeaileTagBo
	tagBo.ParentId = id
	response := tagService.GetTagListByUser(ctx, tagBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func SaveTag(ctx *gin.Context) {
	tagBo := bo.MeaileTagBo{}
	if err := ctx.ShouldBind(&tagBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	tagService := impl.TagServiceImpl{}
	response := tagService.SaveTag(ctx, tagBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func UpdateTag(ctx *gin.Context) {
	tagBo := bo.MeaileTagBo{}
	if err := ctx.ShouldBind(&tagBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	tagService := impl.TagServiceImpl{}
	response := tagService.UpdateTag(ctx, tagBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func DeleteTag(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	tagService := impl.TagServiceImpl{}
	response := tagService.DeleteTag(ctx, id)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
