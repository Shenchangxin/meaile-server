package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	//UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
	UserRouter.POST("login", controller.Login)
	UserRouter.POST("register", controller.Register)
}
