package middleware

import (
	"net/http"

	"github.com/Zzocker/blab/adapters"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core/ports"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/gin-gonic/gin"
)

type oauthMi struct{}

var (
	tokenStore ports.OAuthStore
)

func (o *oauthMi) build(conf config.C) errors.E {
	db, err := adapters.CreateOAuthStore(conf.Core.OAuth.TokenStoreConf)
	if err != nil {
		return err
	}
	tokenStore = db
	return nil
}

func OAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		handleToken(c)
		c.Next()
	}
}

func handleToken(c *gin.Context) {
	tokenID := c.GetHeader("Authorization")
	token, err := tokenStore.Get(c.Request.Context(), tokenID)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if token.Type != model.AccessToken {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("username", token.Username)
}
