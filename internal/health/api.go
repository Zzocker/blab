package health

import (
	"net/http"

	"github.com/Zzocker/blab/pkg/response"
	"github.com/gin-gonic/gin"
)

// RegisterHandlers registers the handlers that perform healthchecks
func RegisterHandlers(r *gin.Engine) {
	r.GET("/ping", check)
}

func check(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusInternalServerError, "not pong")
}
