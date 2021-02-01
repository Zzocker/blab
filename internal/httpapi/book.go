package httpapi

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core"
	"github.com/Zzocker/blab/core/book"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/gin-gonic/gin"
)

type bookAPI struct {
	c core.BookCore
}

func (b *bookAPI) RegisterHandlers(conf config.C, nonAuthR, authR *gin.RouterGroup) errors.E {
	c, err := book.NewBookCore(conf)
	if err != nil {
		return err
	}
	b.c = c

	// Register OAuth endpoints
	// Register NonOauth endpoints

	return nil
}
