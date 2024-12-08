package impl

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/middlewares"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	vo "meaile-server/meaile-user/model/vo"
	"time"
)

type BookServiceImpl struct {
}

func (b *BookServiceImpl) SaveBook(ctx *gin.Context, bo bo.MeaileBookBo) *model.Response {
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
	book := model.MeaileBook{
		BookName:     bo.BookName,
		Image:        bo.Image,
		Introduction: bo.Introduction,
		Favorite:     bo.Favorite,
		Sort:         bo.Sort,
		Status:       bo.Status,
		CreatedBy:    customClaims.UserName,
		CreatedTime:  time.Now(),
		UpdatedBy:    customClaims.UserName,
		UpdatedTime:  time.Now(),
	}
	result := global.DB.Create(&book)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "创建失败",
			Data: err,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "创建成功",
		Data: book,
	}
}

func (b *BookServiceImpl) GetBookListByTagId(ctx *gin.Context, bo bo.BookQueryBo) *model.Response {
	var bookList []vo.MeaileBookVo
	result := global.DB.Table("meaile_book mb").
		Select("mb.*,mt.*").
		Joins("left join meaile_book_tag mbt on mbt.book_id = mb.id").
		Joins("left join meaile_book_tag mbt2 on mbt2.book_id = mb.id").
		Joins("left join meaile_tag mt on mt.id = mbt2.tag_id").
		Where("mbt.tag_id = ?", bo.TagId).Order("mb." + bo.SortField + " " + bo.AscOrDesc).
		Scan(&bookList)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "查询成功",
		Data: bookList,
	}
}
