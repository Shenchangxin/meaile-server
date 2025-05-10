package controller

import (
	"github.com/gin-gonic/gin"
	model "meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
	"strconv"
)

// Login 注册/**
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
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": response.Code,
	//	"msg":  response.Msg,
	//	"data": response.Data,
	//})
	return
}

// Register 注册/**
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
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

// UpdateUserInfo 修改用户信息/**
func UpdateUserInfo(ctx *gin.Context) {
	registerUserBo := model.MeaileUserBo{}
	if err := ctx.ShouldBind(&registerUserBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	userService := impl.UserServiceImpl{}
	response := userService.UpdateUser(ctx, registerUserBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

// GetUserFriendList 获取用户好友以及分组列表/**
func GetUserFriendList(ctx *gin.Context) {
	userService := impl.UserServiceImpl{}
	response := userService.GetUserFriendList(ctx)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

// GetUserInfo 获取用户详细信息/**
func GetUserInfo(ctx *gin.Context) {
	userService := impl.UserServiceImpl{}
	response := userService.GetUserInfo(ctx)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return

}

// AddFriend 添加好友/**
func AddFriend(ctx *gin.Context) {
	addUserFriendBo := model.AddUserFriendBo{}
	if err := ctx.ShouldBind(&addUserFriendBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	userService := impl.UserServiceImpl{}
	response := userService.AddFriend(ctx, addUserFriendBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return

}

// DeleteFriend 删除好友/**
func DeleteFriend(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "参数错误",
			"data": err,
		})
		return
	}
	userService := impl.UserServiceImpl{}
	response := userService.DeleteFriend(ctx, id)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return

}
