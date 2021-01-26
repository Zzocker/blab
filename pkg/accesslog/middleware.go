package accesslog

import (
	"time"

	"github.com/Zzocker/blab/pkg/log"
	"github.com/gin-gonic/gin"
)

// Handler returns a middleware that record an access log message fro every http request
func Handler(l log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.WithFields(l, map[string]interface{}{
			"method":  c.Request.Method,
			"status":  c.Writer.Status(),
			"latency": time.Since(start),
			"path":    c.Request.RequestURI,
		}).Info()
	}
}
