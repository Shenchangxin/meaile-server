package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	//UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
	UserRouter.GET("getUserInfo", middlewares.JWTAuth(), controller.GetUserInfo)
	UserRouter.POST("login", controller.Login)
	UserRouter.POST("register", controller.Register)
}
