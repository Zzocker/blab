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

func (u *userAPI) RegisterHandlers(conf config.C, nonAuthR, authR *gin.RouterGroup) errors.E {
	c, err := user.NewUserCore(conf.Core.User)
	if err != nil {
		return err
	}
	// u = &userAPI{}
	u.c = c

	/// non auth routers
	nonAuthR.POST("/user/register", u.register)
	/// auth router
	authR.GET("/user/:username", u.get)
	authR.PATCH("/user", u.update)
	authR.DELETE("/user", u.delete)
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
		log.L.Error(err)
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, usr)
}

func (u *userAPI) get(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		log.L.Error("empty username")
		sendMsg(c, http.StatusBadRequest, "empty username")
		return
	}
	usr, err := u.c.Get(c.Request.Context(), username)
	if err != nil {
		log.L.Error(err)
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, usr)
}

func (u *userAPI) update(c *gin.Context) {
	username := c.GetString("username")
	if username == "" {
		log.L.Error("empty username")
		sendMsg(c, http.StatusBadRequest, "empty username")
		return
	}
	usr, err := u.c.Update(c.Request.Context(), username, c.Request.Body)
	if err != nil {
		log.L.Error(err)
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, usr)
}

func (u *userAPI) delete(c *gin.Context) {
	username := c.GetString("username")
	if username == "" {
		log.L.Error("empty username")
		sendMsg(c, http.StatusBadRequest, "empty username")
		return
	}
	err := u.c.Delete(c.Request.Context(), username)
	if err != nil {
		log.L.Error(err)
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, nil)
}
