package service

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
)

type BookService interface {
	SaveBook(ctx *gin.Context, bo bo.MeaileBookBo) *model.Response
}
