package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitTagRouter(Router *gin.RouterGroup) {
	TagRouter := Router.Group("tag")
	TagRouter.GET("getTagListByParentId", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.TagListByParentId)
	TagRouter.GET("getTagListByUser", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.TagListByUser)
	TagRouter.POST("saveTag", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.SaveTag)
	TagRouter.PUT("updateTag", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.UpdateTag)
	TagRouter.DELETE("deleteTag/:id", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.DeleteTag)
}
