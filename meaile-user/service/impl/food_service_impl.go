package impl

import (
	"errors"
	"github.com/elliotchance/pie/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/middlewares"
	"meaile-server/meaile-user/model"
	model2 "meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	vo "meaile-server/meaile-user/model/vo"
	"net/http"
	"strings"
	"time"
)

type FoodServiceImpl struct {
}

func (f *FoodServiceImpl) SaveFood(ctx *gin.Context, bo bo.MeaileFoodBo) *model.Response {
	token := ctx.Request.Header.Get("X-Token")
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
	claims, _ := ctx.Get("claims")
	customClaims := claims.(*model.CustomClaims)
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
	claims, _ := ctx.Get("claims")
	customClaims := claims.(*model.CustomClaims)
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
func (f *FoodServiceImpl) GetMyFoodList(ctx *gin.Context, query bo.FoodQuery) *model.Response {
	claims, _ := ctx.Get("claims")
	customClaims := claims.(*model.CustomClaims)
	var foods []model.MeaileFood
	offset := (query.PageNum - 1) * query.PageSize
	db := global.DB.Offset(offset).Limit(query.PageSize)
	db.Where("created_by = ?", customClaims.UserName)
	if query.FoodName != "" {
		db.Where("food_name like %?%", query.FoodName)
	}
	if query.TagId != "" {
		db.Joins("inner join meaile_food_tag mft on mft.tag_id = ?", query.TagId)
	}
	result := db.Order("favorite DESC").Find(&foods)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: foods,
	}
}
func (f *FoodServiceImpl) GetFoodList(ctx *gin.Context, query bo.FoodQuery) *model.Response {
	var foods []vo.MeaileFoodVo
	offset := (query.PageNum - 1) * query.PageSize
	db := global.DB.Offset(offset).Limit(query.PageSize)
	if query.FoodName != "" {
		db.Where("food_name like %?%", query.FoodName)
	}
	if query.TagId != "" {
		db.Joins("inner join meaile_food_tag mft on mft.tag_id = ?", query.TagId)
	}
	result := db.Order("favorite DESC").Find(&foods)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	var creators []string
	for _, food := range foods {
		creators = append(creators, food.CreatedBy)
	}
	creatorsStr := strings.Join(creators, ", ")
	var users []vo.MeaileUserVo
	result = global.DB.Where("user_name in (?)", creatorsStr).Find(&users)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	for _, food := range foods {
		for _, user := range users {
			if user.UserName == food.CreatedBy {
				food.Creator = user
			}
		}
	}
	return &model.Response{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: foods,
	}
}
func (f *FoodServiceImpl) GetFollowFoodList(ctx *gin.Context, query bo.FoodQuery) *model.Response {
	claims, _ := ctx.Get("claims")
	customClaims := claims.(*model.CustomClaims)
	var foods []vo.MeaileFoodVo
	offset := (query.PageNum - 1) * query.PageSize
	var userFollowList []vo.MeaileUserFollowVo
	result := global.DB.Where("user_name = ?", customClaims.UserName).Find(&userFollowList)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	// 提取 UserIdFriend 属性
	userIds := pie.Map(userFollowList, func(f vo.MeaileUserFollowVo) string {
		return f.FollowUserName
	})
	db := global.DB.Offset(offset).Limit(query.PageSize)

	result = db.Where("created_by in (?)", userIds).Order("favorite DESC").Find(&foods)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	var creators []string
	for _, food := range foods {
		creators = append(creators, food.CreatedBy)
	}
	//creatorsStr := strings.Join(creators, ", ")
	var users []vo.MeaileUserVo
	result = global.DB.Where("user_name in (?)", creators).Find(&users)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	for _, food := range foods {
		for _, user := range users {
			if user.UserName == food.CreatedBy {
				food.Creator = user
			}
		}
	}
	var imageOssIds []string
	for _, food := range foods {
		imageOssIds = append(imageOssIds, food.Image)
	}
	var ossList []model.MeaileOss
	result = global.DB.Where("oss_id in (?)", imageOssIds).Find(&ossList)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	for i, foodVo := range foods {
		for _, oss := range ossList {
			if oss.OssId == foodVo.Image {
				fileUrl, _ := global.MinioClient.GetPresignedGetObject(global.ServerConfig.MinioConfig.BucketName, oss.OssId+oss.Suffix, 24*time.Hour)
				oss.FileUrl = fileUrl
				foods[i].ImageOssObj = oss
				break
			}
		}
	}
	return &model.Response{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: foods,
	}
}
func (f *FoodServiceImpl) GetRecommendFoodList(ctx *gin.Context, query bo.FoodQuery) *model.Response {
	var foods []vo.MeaileFoodVo
	offset := (query.PageNum - 1) * query.PageSize
	db := global.DB.Offset(offset).Limit(query.PageSize)

	result := db.Order("favorite DESC").Find(&foods)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	var creators []string
	for _, food := range foods {
		creators = append(creators, food.CreatedBy)
	}
	//creatorsStr := strings.Join(creators, ", ")
	var users []vo.MeaileUserVo
	result = global.DB.Where("user_name in (?)", creators).Find(&users)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	for _, food := range foods {
		for _, user := range users {
			if user.UserName == food.CreatedBy {
				food.Creator = user
			}
		}
	}
	var imageOssIds []string
	for _, food := range foods {
		imageOssIds = append(imageOssIds, food.Image)
	}
	var ossList []model.MeaileOss
	result = global.DB.Where("oss_id in (?)", imageOssIds).Find(&ossList)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	for i, foodVo := range foods {
		for _, oss := range ossList {
			if oss.OssId == foodVo.Image {
				fileUrl, _ := global.MinioClient.GetPresignedGetObject(global.ServerConfig.MinioConfig.BucketName, oss.OssId+oss.Suffix, 24*time.Hour)
				oss.FileUrl = fileUrl
				foods[i].ImageOssObj = oss
				break
			}
		}
	}
	return &model.Response{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: foods,
	}
}
func (f *FoodServiceImpl) GetFoodInfo(ctx *gin.Context, id int64) *model.Response {
	var foodInfo vo.MeaileFoodVo
	result := global.DB.Where("id = ?", id).First(&foodInfo)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	if result.RowsAffected == 0 {
		return &model.Response{
			Code: model.SUCCESS,
			Msg:  "未找到该菜品信息",
			Data: nil,
		}
	}
	var creator vo.MeaileUserVo
	result = global.DB.Where("user_name = ?", foodInfo.CreatedBy).First(&creator)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	foodInfo.Creator = creator
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "查询成功",
		Data: foodInfo,
	}
}
