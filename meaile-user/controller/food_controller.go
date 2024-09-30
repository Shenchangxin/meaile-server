package controller

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model/bo"
	"meaile-server/meaile-user/service/impl"
	"net/http"
	"strconv"
	"strings"
)

func SaveFood(ctx *gin.Context) {
	foodBo := model.MeaileFoodBo{}
	if err := ctx.ShouldBind(&foodBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	foodService := impl.FoodServiceImpl{}
	response := foodService.SaveFood(ctx, foodBo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}

func DeleteFood(ctx *gin.Context) {
	ids := ctx.Query("ids")
	idStrings := strings.Split(ids, ",")
	idNumbers := make([]int64, 0, len(idStrings))
	for _, idStr := range idStrings {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "参数转换错误: " + idStr})
			return
		}
		idNumbers = append(idNumbers, id)
	}
	foodService := impl.FoodServiceImpl{}
	response := foodService.DeleteFood(ctx, idNumbers)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func UpdateFood(ctx *gin.Context) {
	foodBo := model.MeaileFoodBo{}
	if err := ctx.ShouldBind(&foodBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	foodService := impl.FoodServiceImpl{}
	response := foodService.UpdateFood(ctx, foodBo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
