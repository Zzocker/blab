package httpapi

import (
	"net/http"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core"
	"github.com/Zzocker/blab/core/user"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/log"
	"github.com/gin-gonic/gin"
)

type userAPI struct {
	c core.UserCore
}

func (u *userAPI) RegisterHandlers(conf config.Core, nonAuthR, authR *gin.RouterGroup) errors.E {
	c, err := user.NewUserCore(conf.User)
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
	var in user.Register
	jErr := c.ShouldBindJSON(&in)
	if jErr != nil {
		log.L.Error(jErr)
		sendMsg(c, http.StatusBadRequest, "wrong json request")
		return
	}
	pass := c.GetHeader("Secret")
	if pass == "" {
		log.L.Error("empty password")
		sendMsg(c, http.StatusBadRequest, "empty password")
		return
	}
	in.Password = pass
	usr, err := u.c.Register(c.Request.Context(), in)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, usr)
}
