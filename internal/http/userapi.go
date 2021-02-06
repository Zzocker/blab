package http

import (
	"fmt"
	"net/http"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core"
	"github.com/Zzocker/blab/internal/logger"
	scode "github.com/Zzocker/blab/pkg/code"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/util"
	"github.com/gin-gonic/gin"
)

const (
	userRouterLoggerPrefix = "[http-user] %v"
)

type userRouter struct {
	c core.UserCore
}

func (u *userRouter) RegisterHandlers(conf config.ApplicationConf, oauth, noOauth *gin.RouterGroup) errors.E {
	logger.L.Info(userRouterLoggerPrefix, "registering handlers")
	logger.L.Info(userRouterLoggerPrefix, "getting user core")
	c := core.GetUserCore()
	u.c = c
	// register non-oauth handlers
	noOauth.POST("/register", u.register)
	// register oauth handlers
	oauth.GET("/get_user/:id", u.getUser)
	oauth.PUT("/", u.updateUser)
	oauth.DELETE("/", u.deleteUser)
	logger.L.Info(userRouterLoggerPrefix, "handler registered")
	return nil
}

func (u *userRouter) register(c *gin.Context) {
	logger.L.Info(userRouterLoggerPrefix, "received register request")
	pass := c.GetHeader("password") //
	var in core.RegisterUserInput
	if jErr := c.ShouldBindJSON(&in); jErr != nil {
		logger.L.Error(userRouterLoggerPrefix, fmt.Sprintf("failed to decode register json request : %v", jErr))
		sendMsg(c, http.StatusBadRequest, "invalid json input")
		return
	}
	err := u.c.Register(c.Request.Context(), in, pass)
	if err != nil {
		logger.L.Error(userRouterLoggerPrefix, err)
		sendMsg(c, scode.ToHTTP(err.GetCode()), err.Error())
		return
	}
	sendData(c, http.StatusCreated, nil)
}

func (u *userRouter) getUser(c *gin.Context) {
	logger.L.Info(userRouterLoggerPrefix, "received get user request")
	usr, err := u.c.GetUser(c.Request.Context(), c.Param("id"))
	if err != nil {
		logger.L.Error(userRouterLoggerPrefix, err)
		sendMsg(c, scode.ToHTTP(err.GetCode()), err.Error())
		return
	}
	sendData(c, http.StatusOK, usr)
}

func (u *userRouter) updateUser(c *gin.Context) {
	logger.L.Info(userRouterLoggerPrefix, "received update user request")
	usr, err := u.c.Update(util.WrapCtx(c.Request.Context(), util.UsernameCtxKey, c.GetHeader("owner")), c.Request.Body) // remove this use Oauth TOTO
	if err != nil {
		logger.L.Error(userRouterLoggerPrefix, err)
		sendMsg(c, scode.ToHTTP(err.GetCode()), err.Error())
		return
	}
	sendData(c, http.StatusOK, usr)
}

func (u *userRouter) deleteUser(c *gin.Context) {
	logger.L.Info(userRouterLoggerPrefix, "received delete user request")
	err := u.c.Delete(util.WrapCtx(c.Request.Context(), util.UsernameCtxKey, c.GetHeader("owner")))
	if err != nil {
		logger.L.Error(userRouterLoggerPrefix, err)
		sendMsg(c, scode.ToHTTP(err.GetCode()), err.Error())
		return
	}
	sendData(c, http.StatusOK, nil)
}
