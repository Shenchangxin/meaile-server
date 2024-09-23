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
	UserRouter.PUT("updateUserInfo", middlewares.JWTAuth(), controller.UpdateUserInfo)
	UserRouter.PUT("getFriendList", middlewares.JWTAuth(), controller.GetUserFriendList)
	UserRouter.POST("addUserFriends", middlewares.JWTAuth(), controller.AddFriend)
	UserRouter.POST("login", controller.Login)
	UserRouter.POST("register", controller.Register)
}
