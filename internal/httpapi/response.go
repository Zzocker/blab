package httpapi

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Data   interface{} `json:"data"`
	Status status      `json:"status"`
}
type status struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func sendData(c *gin.Context, code int, data interface{}) {
	res := response{
		Data: data,
		Status: status{
			Code: code,
		},
	}
	c.JSON(res.Status.Code, res)
}

func sendMsg(c *gin.Context, code int, msg string) {
	res := response{
		Status: status{
			Code:    code,
			Message: msg,
		},
	}
	c.JSON(res.Status.Code, res)
}
