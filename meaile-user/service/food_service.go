package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
)

type FoodService interface {
	SaveFood(ctx *gin.Context, bo bo.MeaileFoodBo) *model.Response
	DeleteFood(ctx *gin.Context, ids []int64) *model.Response
	UpdateFood(ctx *gin.Context, bo bo.MeaileFoodBo) *model.Response
}