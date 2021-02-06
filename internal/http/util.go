package http

import "github.com/gin-gonic/gin"

type res struct {
	Data interface{} `json:"data"`
	Code code        `json:"code"`
}

type code struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

func sendMsg(c *gin.Context, status int, msg string) {
	response := res{
		Code: code{
			Status:  status,
			Message: msg,
		},
	}
	c.JSON(status, response)
}

func sendData(c *gin.Context, status int, data interface{}) {
	response := res{
		Data: data,
		Code: code{
			Status: status,
		},
	}
	c.JSON(status, response)
}
