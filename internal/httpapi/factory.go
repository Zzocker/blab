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
		&oauthAPI{},
	}
)

// HTTPRouter :
type httpRouter interface {
	RegisterHandlers(conf config.C, nonAuthR, authR *gin.RouterGroup) errors.E
}

// BuildAllRouter : build all routers
func BuildAllRouter(conf config.C, nonAuthR, authR *gin.RouterGroup) {
	for i := range f {
		err := f[i].RegisterHandlers(conf, nonAuthR, authR)
		if err != nil {
			log.L.Fatal("Failed to register", err)
		}
	}
	log.L.Info("All endpoints registered")
}