package http

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/pkg/errors"
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
	// register oauth handlers
	noOauth.POST("/register", u.register)
	logger.L.Info(userRouterLoggerPrefix, "handler registered")
	return nil
}

func (u *userRouter) register(c *gin.Context) {

}
