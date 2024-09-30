package impl

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/middlewares"
	"meaile-server/meaile-user/model"
	model2 "meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	"net/http"
	"time"
)

type FoodServiceImpl struct {
}

func (f *FoodServiceImpl) SaveFood(ctx *gin.Context, bo bo.MeaileFoodBo) *model.Response {
	token := ctx.Request.Header.Get("x-token")
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err,
		}
	}
	food := model2.MeaileFood{
		FoodName:     bo.FoodName,
		Image:        bo.Image,
		Introduction: bo.Introduction,
		Favorite:     bo.Favorite,
		CreatedTime:  time.Now(),
		CreatedBy:    customClaims.UserName,
		UpdatedTime:  time.Now(),
		UpdatedBy:    customClaims.UserName,
	}
	tx := global.DB.Begin()
	result := tx.Create(&food)
	if result.Error != nil {
		return &model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "保存失败",
			Data: result.Error,
		}
	}
	var tags []model2.MeaileFoodTag
	for _, tag := range bo.Tags {
		tagInsert := model2.MeaileFoodTag{
			FoodId:      food.Id,
			Tag:         tag,
			Status:      "0",
			CreatedBy:   customClaims.UserName,
			CreatedTime: time.Now(),
		}
		tags = append(tags, tagInsert)
	}
	result = tx.Create(&tags)
	if result.Error != nil {
		tx.Rollback()
		return &model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "保存失败",
			Data: result.Error,
		}
	}
	tx.Commit()
	return &model.Response{
		Code: http.StatusOK,
		Msg:  "保存成功",
		Data: bo,
	}
}
func (f *FoodServiceImpl) DeleteFood(ctx *gin.Context, ids []int64) *model.Response {
	token := ctx.Request.Header.Get("x-token")
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err,
		}
	}
	var foods []model2.MeaileFood
	result := global.DB.Where("id in ? and created_by = ?", ids, customClaims.UserName).Find(&foods)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "操作失败",
			Data: result.Error,
		}
	}
	if int64(len(ids)) > result.RowsAffected {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "参数错误",
			Data: nil,
		}
	}
	tx := global.DB.Begin()

	result = tx.Delete(&model2.MeaileFood{}, ids)
	if result.Error != nil {
		tx.Rollback()
		return &model.Response{
			Code: model.FAILED,
			Msg:  "删除失败",
			Data: result.Error,
		}
	}
	result = tx.Where("food_id in ?", ids).Delete(&model2.MeaileFood{})
	if result.Error != nil {
		tx.Rollback()
		return &model.Response{
			Code: model.FAILED,
			Msg:  "删除失败",
			Data: result.Error,
		}
	}
	tx.Commit()
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: ids,
	}
}
func (f *FoodServiceImpl) UpdateFood(ctx *gin.Context, bo bo.MeaileFoodBo) *model.Response {
	token := ctx.Request.Header.Get("x-token")
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err,
		}
	}
	food := model2.MeaileFood{
		FoodName:     bo.FoodName,
		Image:        bo.Image,
		Introduction: bo.Introduction,
		Favorite:     bo.Favorite,
		UpdatedTime:  time.Now(),
		UpdatedBy:    customClaims.UserName,
	}
	tx := global.DB.Begin()
	result := tx.Model(&model2.MeaileFood{}).Where("id = ? and create_by = ?", bo.Id, customClaims.UserName).Omit("id").Updates(food)
	if result.Error != nil {
		return &model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "更新失败",
			Data: result.Error,
		}
	}
	if result.RowsAffected != 1 {
		tx.Rollback()
		return &model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "更新失败",
			Data: result.Error,
		}
	}
	result = tx.Where("food_id = ?", bo.Id).Delete(&model2.MeaileFoodTag{})
	if result.Error != nil {
		tx.Rollback()
		return &model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "更新失败",
			Data: result.Error,
		}
	}
	var tags []model2.MeaileFoodTag
	for _, tag := range bo.Tags {
		tagInsert := model2.MeaileFoodTag{
			FoodId:      food.Id,
			Tag:         tag,
			Status:      "0",
			CreatedBy:   customClaims.UserName,
			CreatedTime: time.Now(),
		}
		tags = append(tags, tagInsert)
	}
	result = tx.Create(&tags)
	if result.Error != nil {
		tx.Rollback()
		return &model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "保存失败",
			Data: result.Error,
		}
	}
	tx.Commit()
	return &model.Response{
		Code: http.StatusOK,
		Msg:  "保存成功",
		Data: bo,
	}
}
