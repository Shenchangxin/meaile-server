package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS int = 200 //操作成功
	FAILED  int = 500 //操作失败
)

// 请求成功的时候 使用该方法返回信息
func Success(ctx *gin.Context, code int, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"smg":  "成功",
		"data": v,
	})
}

// 请求失败的时候, 使用该方法返回信息
func Failed(ctx *gin.Context, code int, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"data": nil,
		"msg":  v,
	})
}
