package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitFoodRouter(Router *gin.RouterGroup) {
	FoodRouter := Router.Group("food")
	FoodRouter.POST("saveFood", middlewares.JWTAuth(), controller.SaveFood)
	FoodRouter.POST("updateFood", middlewares.JWTAuth(), controller.UpdateFood)
	FoodRouter.GET("deleteFood", middlewares.JWTAuth(), controller.DeleteFood)
	//FoodRouter.GET("getFoodList", middlewares.JWTAuth(), controller.FoodList)
	//FoodRouter.GET("getFoodInfo/:id", middlewares.JWTAuth(), controller.FoodInfo)
}
