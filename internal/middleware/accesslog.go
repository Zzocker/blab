package middleware

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/internal/util"
	"github.com/gin-gonic/gin"
)

const (
	accessPkg = "accesslog"
)

var (
	requestCounter int64 = 0
)

func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := atomic.AddInt64(&requestCounter, 1)
		start := time.Now()
		c.Set(fmt.Sprint(util.RequestIDKey), reqID)
		c.Next()
		logger.L.Debugf(reqID, accessPkg,
			"Latency=%v Method=%s Path=%s Status=%d WriteByteSize=%v",
			time.Since(start),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			c.Writer.Size(),
		)
	}
}
