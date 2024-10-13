package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	//UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
	UserRouter.GET("getUserInfo", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetUserInfo)
	UserRouter.PUT("updateUserInfo", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.UpdateUserInfo)
	UserRouter.GET("getFriendList", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetUserFriendList)
	UserRouter.POST("addUserFriends", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.AddFriend)
	UserRouter.GET("deleteFriend", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.DeleteFriend)
	UserRouter.POST("login", controller.Login)
	UserRouter.POST("register", controller.Register)
}
