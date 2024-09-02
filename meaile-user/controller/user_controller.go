package controller

import (
	"github.com/gin-gonic/gin"
	model "meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
)

func Login(ctx *gin.Context) {
	loginForm := model.LoginForm{}
	if err := ctx.ShouldBind(&loginForm); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
	}
	userService := impl.UserServiceImpl{}
	response := userService.Login(ctx, loginForm)
	if response.Code != 200 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": response.Code,
			"msg":  response.Msg,
			"data": response.Data,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": response.Code,
			"msg":  response.Msg,
			"data": response.Data,
		})
	}
}
