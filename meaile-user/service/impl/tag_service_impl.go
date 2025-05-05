package impl

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/middlewares"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	"time"
)

type TagServiceImpl struct {
}

func (t *TagServiceImpl) GetTagListByParentId(ctx *gin.Context, parentId int64) *model.Response {
	var tags []model.MeaileTag
	result := global.DB.Where("parent_id = ?", parentId).Find(&tags)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询分组列表失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: tags,
	}
}

func (*TagServiceImpl) GetTagListByUser(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response {
	var tagList []model.MeaileTag
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
	db := global.DB.Where("created_by = ?", customClaims.UserName)
	if tagBo.ParentId != -1 {
		db.Where("parent_id = ?", tagBo.ParentId)
	}
	result := db.Find(&tagList)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询分组列表失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: tagList,
	}
}
func (*TagServiceImpl) SaveTag(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response {
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
	var tag model.MeaileTag
	result := global.DB.Where("created_by = ? and tagName = ?", customClaims.UserName, tagBo.TagName).Find(&tag)
	if result.RowsAffected == 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "该标签已存在",
			Data: err,
		}
	}
	tag.TagName = tagBo.TagName
	tag.CreatedBy = customClaims.UserName
	tag.ParentId = tagBo.ParentId
	tag.CreatedTime = time.Now()
	tag.UpdatedBy = customClaims.UserName
	tag.UpdatedTime = time.Now()
	tag.Status = "0"
	result = global.DB.Create(&tag)
	if result.Error == nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "新建标签信息失败",
			Data: err,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "新建标签信息成功",
		Data: tag,
	}
}

func (*TagServiceImpl) UpdateTag(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response {
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
	var tag model.MeaileTag
	result := global.DB.Where("created_by = ? and id = ?", customClaims.UserName, tagBo.Id).Find(&tag)
	if result.RowsAffected != 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "未找到该标签信息",
			Data: err,
		}
	}
	tag.Id = tagBo.Id
	tag.TagName = tagBo.TagName
	tag.UpdatedBy = customClaims.UserName
	tag.UpdatedTime = time.Now()
	tag.Status = tagBo.Status
	result = global.DB.Updates(&tag)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "更新标签信息失败",
			Data: err,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "更新标签信息成功",
		Data: tag,
	}
}

func (*TagServiceImpl) DeleteTag(ctx *gin.Context, id int64) *model.Response {
	var tag model.MeaileTag
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
	result := global.DB.Where("created_by = ? and id = ?", customClaims.UserName, id).Find(&tag)
	if result.RowsAffected != 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "未找到要删除的标签信息",
			Data: err,
		}
	}
	result = global.DB.Where("id = ?", id).Delete(&tag)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "删除失败",
			Data: err,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "删除成功",
		Data: tag,
	}
}

//func (*TagServiceImpl) GetTagInfo(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response {
//	return nil
//}
