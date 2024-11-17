package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitTagRouter(Router *gin.RouterGroup) {
	TagRouter := Router.Group("tag")
	TagRouter.GET("getTagListByParentId", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.TagListByParentId)
}
