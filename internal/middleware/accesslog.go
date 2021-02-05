package middleware

import (
	"fmt"
	"time"

	"github.com/Zzocker/blab/internal/logger"
	"github.com/gin-gonic/gin"
)

const (
	accessloggerPrefix = "[middleware-access] %v"
)

// AccessLog : logger middleware called first
// for logging request and responses
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.L.Info(accessloggerPrefix,
			fmt.Sprintf("Latency=%v Method=%s Path=%s Status=%d WriteByteSize=%v",
				time.Since(start),
				c.Request.Method,
				c.Request.URL.Path,
				c.Writer.Status(),
				c.Writer.Size()))
	}
}
