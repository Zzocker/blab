package health

import (
	"net/http"

	"github.com/Zzocker/blab/pkg/response"
	"github.com/gin-gonic/gin"
)

// RegisterHandlers registers the handlers that perform healthchecks
func RegisterHandlers(r *gin.Engine) {
	r.GET("/healthcheck", check)
}

func check(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusOK, "ok")
}
