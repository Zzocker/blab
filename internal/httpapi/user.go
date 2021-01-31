package httpapi

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/gin-gonic/gin"
)

type userAPI struct {
	c core.UserCore
}

func (u *userAPI) RegisterHandlers(conf config.Core, nonAuthR, authR *gin.RouterGroup) errors.E {
	c, err := core.NewUserCore(conf.User)
	if err != nil {
		return err
	}
	// u = &userAPI{}
	u.c = c
	/// non auth routers

	nonAuthR.POST("/register", u.register)
	/// auth router

	return nil
}

func (u *userAPI) register(c *gin.Context) {
	defer sendMsg(c, errors.CodeInternalErr, "endpoint not implemented")
}
