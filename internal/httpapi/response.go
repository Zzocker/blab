package httpapi

import (
	"net/http"

	"github.com/Zzocker/blab/pkg/errors"
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
			Code: getHTTPCode(code),
		},
	}
	c.JSON(res.Status.Code, res)
}

func sendMsg(c *gin.Context, code int, msg string) {
	res := response{
		Status: status{
			Code:    getHTTPCode(code),
			Message: msg,
		},
	}
	c.JSON(res.Status.Code, res)
}

func getHTTPCode(code int) int {
	httpCode := http.StatusOK
	switch code {
	case errors.CodeNotFound:
		httpCode = http.StatusNotFound
	case errors.CodeAlreadyExists:
		httpCode = http.StatusConflict
	case errors.CodeInvalidArgument:
		httpCode = http.StatusBadRequest
	case errors.CodeInternalErr:
		httpCode = http.StatusInternalServerError
	default:
		httpCode = http.StatusNotImplemented
	}
	return httpCode
}
