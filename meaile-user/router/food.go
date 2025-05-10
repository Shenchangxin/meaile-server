package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitFoodRouter(Router *gin.RouterGroup) {
	FoodRouter := Router.Group("food")
	FoodRouter.POST("saveFood", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.SaveFood)
	FoodRouter.POST("updateFood", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.UpdateFood)
	FoodRouter.GET("deleteFood", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.DeleteFood)
	FoodRouter.GET("getMyFoodList", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetMyFoods)
	FoodRouter.GET("getFoodList", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetFoods)
	FoodRouter.GET("getFollowFoodList", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetFollowFoods)
	FoodRouter.GET("getRecommendFoodList", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetRecommendFoods)
	FoodRouter.GET("getFoodInfo/:id", middlewares.JWTAuth(), controller.GetFoodInfo)
}
