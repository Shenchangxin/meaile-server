package router

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/controller"
	"meaile-server/meaile-user/middlewares"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	CommentRouter := Router.Group("comment")
	CommentRouter.GET("getCommentList", middlewares.JWTAuth(), middlewares.LogMiddleware(), controller.GetComments)
}
