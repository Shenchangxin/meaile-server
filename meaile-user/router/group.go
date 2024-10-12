package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitGroupRouter(Router *gin.RouterGroup) {
	GroupRouter := Router.Group("group")
	GroupRouter.POST("saveGroup", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.SaveGroup)
	GroupRouter.POST("updateGroup", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.UpdateGroup)
	GroupRouter.GET("deleteGroup", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.DeleteGroup)
	GroupRouter.GET("getGroupList", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GroupList)
	GroupRouter.GET("getGroupInfo/:id", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GroupInfo)
}
