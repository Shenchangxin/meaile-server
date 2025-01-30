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
func (b *BookServiceImpl) UpdateBook(ctx *gin.Context, bo bo.MeaileBookBo) *model.Response {
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
	var book model.MeaileBook
	result := global.DB.Where("id = ? and created_by = ?", bo.Id, customClaims.UserName).Find(&book)
	if result.Error != nil || result.RowsAffected != 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "未找到食谱信息",
			Data: err,
		}
	}
	book.BookName = bo.BookName
	book.UpdatedBy = customClaims.UserName
	book.UpdatedTime = time.Now()
	book.Status = bo.Status
	book.Introduction = bo.Introduction
	book.Image = bo.Image
	book.Sort = bo.Sort
	result = global.DB.Updates(&book)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "更新食谱信息失败",
			Data: err,
		}
	}

	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "更新食谱信息成功",
		Data: book,
	}
}
func (b *BookServiceImpl) GetBookListByTagId(ctx *gin.Context, bo bo.BookQueryBo) *model.Response {
	var bookList []vo.MeaileBookVo
	var bookIds []int64
	result := global.DB.Table("meaile_book_tag").Where("tag_id = ?", bo.TagId).Pluck("book_id", &bookIds)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	result = global.DB.Preload("TagList").Where("id in (?)", bookIds).Find(&bookList)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	var imageOssIds []string
	for _, book := range bookList {
		imageOssIds = append(imageOssIds, book.Image)
	}
	var ossList []model.MeaileOss
	result = global.DB.Where("oss_id in (?)", imageOssIds).Find(&ossList)
	//result = global.DB.Table("meaile_book mb").
	//	Select("mb.*,mt.*").
	//	Joins("left join meaile_book_tag mbt on mbt.book_id = mb.id").
	//	Joins("left join meaile_book_tag mbt2 on mbt2.book_id = mb.id").
	//	Joins("left join meaile_tag mt on mt.id = mbt2.tag_id").
	//	Where("mbt.tag_id = ?", bo.TagId).Order("mb." + bo.SortField + " " + bo.AscOrDesc).
	//	Scan(&bookList)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	for i, bookVo := range bookList {
		for _, oss := range ossList {
			if oss.OssId == bookVo.Image {
				fileUrl, _ := global.MinioClient.GetPresignedGetObject(global.ServerConfig.MinioConfig.BucketName, oss.OssId+oss.Suffix, 24*time.Hour)
				oss.FileUrl = fileUrl
				bookList[i].ImageOssObj = oss
				break
			}
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "查询成功",
		Data: bookList,
	}
}

func (b *BookServiceImpl) DeleteBook(ctx *gin.Context, id int64) *model.Response {

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
	var book model.MeaileBook
	result := global.DB.Where("id = ? and created_by = ?", id, customClaims.UserName).Find(&book)
	if result.Error != nil || result.RowsAffected != 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "未找到食谱信息",
			Data: result.Error,
		}
	}
	result = global.DB.Where("id = ?", id).Delete(&book)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "删除失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "删除成功",
		Data: book,
	}
}

func (b *BookServiceImpl) GetBookInfo(ctx *gin.Context, id int64) *model.Response {
	var book model.MeaileBook
	result := global.DB.Where("id = ? ", id).First(&book)
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
		Data: book,
	}
}

func (b *BookServiceImpl) GetMyBooks(ctx *gin.Context) *model.Response {
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
	var myBooks []vo.MeaileBookVo
	result := global.DB.Where("created_by = ?", customClaims.UserName).Find(&myBooks)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: result.Error,
		}
	}
	var imageOssIds []string
	for _, book := range myBooks {
		imageOssIds = append(imageOssIds, book.Image)
	}
	var ossList []model.MeaileOss
	result = global.DB.Where("oss_id in (?)", imageOssIds).Find(&ossList)
	for i, bookVo := range myBooks {
		for _, oss := range ossList {
			if oss.OssId == bookVo.Image {
				fileUrl, _ := global.MinioClient.GetPresignedGetObject(global.ServerConfig.MinioConfig.BucketName, oss.OssId+oss.Suffix, 24*time.Hour)
				oss.FileUrl = fileUrl
				myBooks[i].ImageOssObj = oss
				break
			}
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "查询成功",
		Data: myBooks,
	}
}
