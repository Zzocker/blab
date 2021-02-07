package http

import (
	"fmt"
	"net/http"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/internal/util"
	"github.com/gin-gonic/gin"
)

const (
	healthPkgName = "http-health"
)

type healthAPI struct{}

func (h *healthAPI) registerHandlers(conf *config.ApplicationConf, public, private *gin.RouterGroup) error {
	logger.L.Infof(-1, healthPkgName, "registering health handlers")
	public.GET("/check", h.check)
	logger.L.Infof(-1, healthPkgName, "all health handlers successfully registered")
	return nil
}

func (h *healthAPI) check(c *gin.Context) {
	reqID := c.GetInt64(fmt.Sprint(util.RequestIDKey))
	logger.L.Debugf(reqID, healthPkgName, "received check request")
	sendData(c, http.StatusOK, nil)
}
