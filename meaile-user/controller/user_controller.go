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
		return
	}
	userService := impl.UserServiceImpl{}
	response := userService.Login(ctx, loginForm)
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func Register(ctx *gin.Context) {
	registerUserBo := model.MeaileUserBo{}
	if err := ctx.ShouldBind(&registerUserBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	userService := impl.UserServiceImpl{}
	response := userService.Register(ctx, registerUserBo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func GetUserInfo(ctx *gin.Context) {
	token := ctx.Request.Header.Get("x-token")
	if token == "" {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"code": 500,
			"msg":  "登录已过期，请重新登录",
			"data": nil,
		})
		return
	}

	userService := impl.UserServiceImpl{}
	userInfo := userService.GetUserInfo(ctx, token)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取用户信息成功",
		"data": userInfo,
	})
	return

}
