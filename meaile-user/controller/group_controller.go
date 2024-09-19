package controller

import (
	"github.com/gin-gonic/gin"
	model "meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
)

func SaveGroup(ctx *gin.Context) {
	groupBo := model.MeaileFriendGroupBo{}
	if err := ctx.ShouldBind(&groupBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	groupService := impl.GroupServiceImpl{}
	response := groupService.SaveGroup(ctx, groupBo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func DeleteGroup(ctx *gin.Context) {
	groupBo := model.DeleteIds{}
	if err := ctx.ShouldBind(&groupBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	groupService := impl.GroupServiceImpl{}
	response := groupService.DeleteGroup(ctx, groupBo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
