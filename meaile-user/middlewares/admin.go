package middlewares

import (
	"github.com/gin-gonic/gin"
	"meaile-web/meaile-user/model"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*model.CustomClaims)

		if currentUser.ID != 1 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()

	}
}
