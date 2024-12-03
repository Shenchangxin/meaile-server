package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	router "meaile-server/meaile-user/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/api/v1")
	zap.S().Info("----配置用户相关URL----")
	router.InitUserRouter(ApiGroup)
	zap.S().Info("----配置分组相关URL----")
	router.InitGroupRouter(ApiGroup)
	zap.S().Info("----配置菜品相关URL----")
	router.InitFoodRouter(ApiGroup)
	zap.S().Info("----配置OSS相关URL----")
	router.InitOssRouter(ApiGroup)
	zap.S().Info("----配置Tag相关URL----")
	router.InitTagRouter(ApiGroup)
	zap.S().Info("----配置Book相关URL----")
	router.InitBookRouter(ApiGroup)
	return Router
}
