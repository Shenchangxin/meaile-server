package controller

import (
	"github.com/gin-gonic/gin"
	model "meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
)

func FollowUser(ctx *gin.Context) {
	followBo := model.MeaileUserFollowBo{}
	if err := ctx.ShouldBind(&followBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	followService := impl.FollowServiceImpl{}
	response := followService.FollowUser(ctx, followBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
