package middleware

import (
	"time"
	"uso.link/internal/utils/log2"

	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latency := endTime.Sub(startTime)

		// 请求方法、路径、状态码等信息
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		// 打印日志
		log2.Logf("[GIN] %3d | %v | %s | %s\n",
			statusCode,
			latency,
			method,
			path,
		)
	}
}
