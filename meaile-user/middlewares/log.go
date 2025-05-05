package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"meaile-server/meaile-user/global"
	"net/http"
	"time"
)

// 记录日志的中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Token")
		myJwt := NewJWT()
		customClaims, err := myJwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "用户信息过期，请重新登录",
				"data": err,
			})
			return
		}
		c.Next() // 处理请求
		// 获取状态码
		status := c.Writer.Status()
		// 将日志写入数据库
		tx := global.DB.Exec("INSERT INTO meaile_log (id,method, path, status, ip,oper_time,oper_user) VALUES (?,?, ?, ?,?, ?,?)", nil, c.Request.Method, c.Request.URL.Path, status, c.ClientIP(), time.Now(), customClaims.UserName)
		if tx.Error != nil {
			zap.S().Errorw("无法插入日志: %v\n", tx.Error)
		}
	}
}
