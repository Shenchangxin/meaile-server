package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
)

type BookService interface {
	SaveBook(ctx *gin.Context, bo bo.MeaileBookBo) *model.Response
	UpdateBook(ctx *gin.Context, bo bo.MeaileBookBo) *model.Response
	GetBookListByTagId(ctx *gin.Context, bo bo.BookQueryBo) *model.Response
	DeleteBook(ctx *gin.Context, id int64) *model.Response
	GetBookInfo(ctx *gin.Context, id int64) *model.Response
}
