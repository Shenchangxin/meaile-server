package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	userRouter "meaile-server/meaile-user/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/v1")
	zap.S().Info("----配置用户相关URL----")
	userRouter.InitUserRouter(ApiGroup)
	userRouter.InitGroupRouter(ApiGroup)
	userRouter.InitFoodRouter(ApiGroup)
	return Router
}
