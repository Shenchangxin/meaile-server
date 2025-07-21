package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitFollowRouter(Router *gin.RouterGroup) {
	FoodRouter := Router.Group("follow")
	FoodRouter.POST("followUser", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.FollowUser)
	FoodRouter.POST("unfollowUser", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.UnfollowUser)

}
