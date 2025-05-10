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
	ctx.JSON(response.Code, gin.H{
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
	ctx.JSON(response.Code, gin.H{
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
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func GetMyFoods(ctx *gin.Context) {
	foodBo := model.FoodQuery{}
	if err := ctx.ShouldBind(&foodBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	foodService := impl.FoodServiceImpl{}
	response := foodService.GetMyFoodList(ctx, foodBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func GetFoods(ctx *gin.Context) {
	foodBo := model.FoodQuery{}
	if err := ctx.ShouldBind(&foodBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	foodService := impl.FoodServiceImpl{}
	response := foodService.GetFoodList(ctx, foodBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func GetFollowFoods(ctx *gin.Context) {
	foodBo := model.FoodQuery{}
	if err := ctx.ShouldBind(&foodBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	foodBo.PageNum = pageNum
	foodBo.PageSize = pageSize
	foodService := impl.FoodServiceImpl{}
	response := foodService.GetFollowFoodList(ctx, foodBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func GetRecommendFoods(ctx *gin.Context) {
	foodBo := model.FoodQuery{}
	if err := ctx.ShouldBind(&foodBo); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "参数错误",
		})
		return
	}
	foodBo.PageNum = pageNum
	foodBo.PageSize = pageSize
	foodService := impl.FoodServiceImpl{}
	response := foodService.GetRecommendFoodList(ctx, foodBo)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
func GetFoodInfo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "参数错误",
			"data": err,
		})
		return
	}
	foodService := impl.FoodServiceImpl{}
	response := foodService.GetFoodInfo(ctx, idInt)
	ctx.JSON(response.Code, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
