package http

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/gin-gonic/gin"
)

const (
	userPkgName = "http-user"
)

type userAPI struct{}

func registerHandlers(conf *config.ApplicationConf, public, private *gin.RouterGroup) error {
	logger.L.Infof(-1, userPkgName, "registering user handlers")
	logger.L.Infof(-1, userPkgName, "all user handlers successfully registered")
	return nil
}
