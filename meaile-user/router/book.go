package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitBookRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("book")
	BookRouter.POST("saveBook", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.SaveBook)
	BookRouter.PUT("updateBook", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.UpdateBook)
	BookRouter.DELETE("deleteBook/:id", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.DeleteBook)
	BookRouter.GET("getBookInfo/:id", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetBookInfo)
	BookRouter.GET("getBookListByTagId", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetBookListByTag)
}
