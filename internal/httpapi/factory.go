package httpapi

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/log"
	"github.com/gin-gonic/gin"
)

var (
	f = []httpRouter{
		&userAPI{},
	}
)

// HTTPRouter :
type httpRouter interface {
	RegisterHandlers(conf config.Core, nonAuthR, authR *gin.RouterGroup) errors.E
}

// BuildAllRouter : build all routers
func BuildAllRouter(conf config.Core, nonAuthR, authR *gin.RouterGroup) errors.E {
	for i := range f {
		err := f[i].RegisterHandlers(conf, nonAuthR, authR)
		if err != nil {
			log.L.Error("Failed to register", err)
			return err
		}
	}
	log.L.Info("All endpoints registered")
	return nil
}