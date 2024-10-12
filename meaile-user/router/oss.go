package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("oss")
	OssRouter.POST("upload", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.Upload)
	OssRouter.POST("download", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.DownLoad)
	OssRouter.GET("getUrl/:id", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetUrl)
}
