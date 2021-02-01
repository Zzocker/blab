package httpapi

import (
	"context"
	"net/http"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core"
	"github.com/Zzocker/blab/core/comment"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/gin-gonic/gin"
)

type commentAPI struct {
	c core.CommentCore
}

func (c *commentAPI) RegisterHandlers(conf config.C, nonAuthR, authR *gin.RouterGroup) errors.E {
	cr, err := comment.NewCommentCore(conf)
	if err != nil {
		return err
	}
	c.c = cr

	// Auth Endpoints
	authR.POST("/comment/user", c.createCommentOnUser) // create comment on user
	// authR.POST("/comment/book")    // create comment on book
	// authR.POST("/comment/comment") // create comment on comment
	// authR.PATCH("/comment/:id")    // update a coment with comment id
	// authR.DELETE("/comment/:id")   // delete a comment
	// authR.GET("/comment/:id")      // get a particular
	// authR.GET("/comment/user")     // get all comment made on a user
	// authR.GET("/comment/book")     // get all comment made on a book
	// authR.GET("/comment/comment")  // get all comment made on a comment

	// NoAuth Endpoints
	return nil
}

func (co *commentAPI) createCommentOnUser(c *gin.Context) {
	var newComment comment.CommentCreateInput
	jErr := c.ShouldBindJSON(&newComment)
	if jErr != nil {
		sendMsg(c, http.StatusBadRequest, "wrong json input")
		return
	}
	// TODO create separate function and also set key a package level constant
	com, err := co.c.CommentOn(context.WithValue(c.Request.Context(), "username", c.GetString("username")), newComment)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, com)
}

func (co *commentAPI) createCommentOnBook(c *gin.Context) {

}

func (co *commentAPI) createCommentOnComment(c *gin.Context) {

}
