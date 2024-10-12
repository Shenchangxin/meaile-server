package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"meaile-server/meaile-user/global"
	"time"
)

// 记录日志的中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 处理请求
		// 获取状态码
		status := c.Writer.Status()
		// 将日志写入数据库
		tx := global.DB.Exec("INSERT INTO meaile_log (method, path, status, timestamp) VALUES (?, ?, ?, ?)", c.Request.Method, c.Request.URL.Path, status, time.Now())
		if tx.Error != nil {
			zap.S().Errorw("无法插入日志: %v\n", tx.Error)
		}
	}
}
