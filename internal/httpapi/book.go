package httpapi

import (
	"net/http"

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
	b.c = c // TODO

	// Register OAuth endpoints
	authR.POST("book/add", b.add)
	authR.GET("book/:isbn", b.get)
	authR.PATCH("/book/:isbn", b.patch)
	authR.DELETE("/book/:isbn", b.delete)
	// Register NonOauth endpoints

	return nil
}

func (b *bookAPI) add(c *gin.Context) {
	var newBook book.BookCreate
	jErr := c.ShouldBindJSON(&newBook)
	if jErr != nil {
		sendMsg(c, http.StatusBadRequest, "invalid json input")
		return
	}
	newBook.Username = c.GetString("username")
	bk, err := b.c.AddBook(c.Request.Context(), newBook)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, bk)
}

func (b *bookAPI) get(c *gin.Context) {
	isbn := c.Param("isbn")
	bk, err := b.c.Get(c.Request.Context(), isbn)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, bk)
}

func (b *bookAPI) patch(c *gin.Context) {
	updated, err := b.c.Update(c.Request.Context(), c.Param("isbn"), c.Request.Body)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, updated)
}

func (b *bookAPI) delete(c *gin.Context) {
	isbn := c.Param("isbn")
	err := b.c.Remove(c.Request.Context(), isbn)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, nil)
}
