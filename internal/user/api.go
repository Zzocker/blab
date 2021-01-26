package user

import (
	"net/http"

	"github.com/Zzocker/blab/pkg/response"
	"github.com/gin-gonic/gin"
)

// RegisterHandlers will register endpoints for user api
func RegisterHandlers(r *gin.Engine) {
	r.POST("/user/register", register)
	r.GET("/user/:username", get)
	r.PUT("/user", put)
	r.PATCH("/user", patch)
	r.DELETE("/user/:username", delete)
}

func register(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusServiceUnavailable, "StatusServiceUnavailable")
}

func get(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusServiceUnavailable, "StatusServiceUnavailable")
}

func put(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusServiceUnavailable, "StatusServiceUnavailable")
}

func patch(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusServiceUnavailable, "StatusServiceUnavailable")
}

func delete(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusServiceUnavailable, "StatusServiceUnavailable")
}
