package controller

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
	"strconv"
)

func SaveGroup(ctx *gin.Context) {
	groupBo := bo.MeaileFriendGroupBo{}
	if err := ctx.ShouldBind(&groupBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	groupService := impl.GroupServiceImpl{}
	response := groupService.SaveGroup(ctx, groupBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func DeleteGroup(ctx *gin.Context) {
	groupBo := bo.DeleteGroupIds{}
	if err := ctx.ShouldBind(&groupBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	groupService := impl.GroupServiceImpl{}
	response := groupService.DeleteGroup(ctx, groupBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func UpdateGroup(ctx *gin.Context) {
	groupBo := bo.MeaileFriendGroupBo{}
	if err := ctx.ShouldBind(&groupBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	groupService := impl.GroupServiceImpl{}
	response := groupService.UpdateGroup(ctx, groupBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func GroupList(ctx *gin.Context) {
	groupService := impl.GroupServiceImpl{}
	claims, _ := ctx.Get("claims")
	customClaims := claims.(*model.CustomClaims)
	response := groupService.GetGroupListByUserId(ctx, int64(customClaims.ID))
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func GroupInfo(ctx *gin.Context) {
	groupService := impl.GroupServiceImpl{}
	groupIdStr := ctx.Param("id")
	// 尝试将字符串id转换为int64
	groupId, err := strconv.ParseInt(groupIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "参数错误",
			"data": err,
		})
		return
	}

	response := groupService.GetGroupById(ctx, groupId)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
