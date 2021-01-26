package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Res represents response object which will be sent to client
type Res interface {
	Send(c *gin.Context)
	SetCode(code int, msg string)
	SetData(code int, data interface{})
}

type res struct {
	Data   interface{} `json:"data"`
	Status status      `json:"status"`
}

type status struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// New Creates a new api response
func New() Res {
	return &res{
		Status: status{
			Code: http.StatusOK,
		},
	}
}

func (r *res) Send(c *gin.Context) {
	c.JSON(r.Status.Code, r)
}

func (r *res) SetCode(code int, msg string) {
	r.Status.Code = code
	r.Status.Message = msg
}

func (r *res) SetData(code int, data interface{}) {
	r.Data = data
	r.Status.Code = code
	r.Status.Message = ""
}
