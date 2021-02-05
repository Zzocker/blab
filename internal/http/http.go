package http

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/gin-gonic/gin"
)

const (
	httpLoggerPrefix = "[http] %v"
)

var (
	// factory stores all router group in map
	// key : groupName ----> value : RouterGroup
	// eg: /user ----> userRouter
	factory = map[string]RouterGroup{
		"user": &userRouter{},
	}
)

// RouterGroup represents similar handlers grouped together
type RouterGroup interface {
	RegisterHandlers(conf config.ApplicationConf, oauth, noOauth *gin.RouterGroup) errors.E
}

func RegisterRouters(conf config.ApplicationConf, oauth, noOauth *gin.RouterGroup) errors.E {
	logger.L.Info(httpLoggerPrefix, "registering routers")
	for routeName := range factory {
		err := factory[routeName].RegisterHandlers(conf, oauth.Group(routeName), noOauth.Group(routeName))
		if err != nil {
			return err
		}
	}
	logger.L.Info(httpLoggerPrefix, "successfully registered all routers")
	return nil
}
