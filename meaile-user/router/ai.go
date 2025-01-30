package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/ai/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitAiRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("ai")
	BookRouter.POST("chatAi", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.HandleChat)
}
