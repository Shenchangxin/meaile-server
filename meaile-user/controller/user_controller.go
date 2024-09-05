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
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
}

func Register(ctx *gin.Context) {
	registerUserBo := model.MeaileUserBo{}
	if err := ctx.ShouldBind(&registerUserBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
	}
	userService := impl.UserServiceImpl{}
	response := userService.Register(ctx, registerUserBo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
}
