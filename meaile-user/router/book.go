package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitBookRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("book")
	BookRouter.POST("saveFood", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.SaveBook)
}
