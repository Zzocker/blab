package httpapi

import (
	"net/http"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core"
	"github.com/Zzocker/blab/core/oauth"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/log"
	"github.com/gin-gonic/gin"
)

type oauthAPI struct {
	c core.OAuthCore
}

func (o *oauthAPI) RegisterHandlers(conf config.C, nonAuthR, authR *gin.RouterGroup) errors.E {
	c, err := oauth.NewOAuthCore(conf)
	if err != nil {
		return err
	}
	o.c = c

	//nonauth endpoints
	nonAuthR.GET("/auth/:username", o.login)
	return nil
}

func (o *oauthAPI) login(c *gin.Context) {
	tokens, err := o.c.Login(c.Request.Context(), c.Param("username"), c.GetHeader("Secret"))
	if err != nil {
		log.L.Error(err)
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, tokens)
}
