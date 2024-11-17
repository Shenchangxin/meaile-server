package impl

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/model"
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

//func (*TagServiceImpl) GetTagListByUser(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response {
//	return nil
//}
//func (*TagServiceImpl) GetTagInfo(ctx *gin.Context, tagBo bo.MeaileTagBo) *model.Response {
//	return nil
//}
