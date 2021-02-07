package http

import (
	"os"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/gin-gonic/gin"
)

const (
	httpPkgName = "http"
)

var (
	factory = map[string]routerBuilder{
		"health": &healthAPI{},
	}
)

type routerBuilder interface {
	registerHandlers(conf *config.ApplicationConf, public, private *gin.RouterGroup) error
}

func RegisterRouters(conf *config.ApplicationConf, public, private *gin.RouterGroup) {
	logger.L.Infof(-1, httpPkgName, "registering routers")
	var err error
	for routerName := range factory {
		err = factory[routerName].registerHandlers(conf, private.Group(routerName), public.Group(routerName))
		if err != nil {
			logger.L.Errorf(-1, httpPkgName, "failed to register %s : %v", routerName, err)
			os.Exit(1)
		}
	}
	logger.L.Infof(-1, httpPkgName, "all routers successfully registered")
}
