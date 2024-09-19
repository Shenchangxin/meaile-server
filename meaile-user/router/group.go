package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitGroupRouter(Router *gin.RouterGroup) {
	GroupRouter := Router.Group("group")
	GroupRouter.POST("saveGroup", middlewares.JWTAuth(), controller.SaveGroup)
	GroupRouter.GET("deleteGroup", middlewares.JWTAuth(), controller.DeleteGroup)
}