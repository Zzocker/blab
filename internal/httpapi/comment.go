package httpapi

import (
	"net/http"
	"strconv"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core"
	"github.com/Zzocker/blab/core/comment"
	"github.com/Zzocker/blab/model"
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
	authR.POST("/comment/user", c.createCommentOnUser)       // create comment on user
	authR.POST("/comment/book", c.createCommentOnBook)       // create comment on book
	authR.POST("/comment/comment", c.createCommentOnComment) // create comment on comment
	authR.PATCH("/comment/:id", c.update)                    // update a coment with comment id
	authR.DELETE("/comment/:id", c.deleteComment)            // delete a comment
	authR.GET("/comment/:id", c.getComment)                  // get a particular
	authR.GET("/comments/user/:username", c.getOnUser)       // get all comment made on a user
	authR.GET("/comments/book/:id", c.getOnBook)             // get all comment made on a book
	authR.GET("/comments/comment/:id", c.getOnComment)       // get all comment made on a comment

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
	com, err := co.c.CommentOn(wrapUsernameOnCtx(c, c.GetString("username")), newComment, model.UserComment)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, com)
}

func (co *commentAPI) createCommentOnBook(c *gin.Context) {
	var newComment comment.CommentCreateInput
	jErr := c.ShouldBindJSON(&newComment)
	if jErr != nil {
		sendMsg(c, http.StatusBadRequest, "wrong json input")
		return
	}
	// TODO create separate function and also set key a package level constant
	com, err := co.c.CommentOn(wrapUsernameOnCtx(c, c.GetString("username")), newComment, model.BookComment)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, com)
}

func (co *commentAPI) createCommentOnComment(c *gin.Context) {
	var newComment comment.CommentCreateInput
	jErr := c.ShouldBindJSON(&newComment)
	if jErr != nil {
		sendMsg(c, http.StatusBadRequest, "wrong json input")
		return
	}
	// TODO create separate function and also set key a package level constant
	com, err := co.c.CommentOn(wrapUsernameOnCtx(c, c.GetString("username")), newComment, model.CommentComment)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, com)
}

func (co *commentAPI) getComment(c *gin.Context) {
	commentID := c.Param("id")
	cmt, err := co.c.GetComment(c.Request.Context(), commentID)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, cmt)
}

func (co *commentAPI) deleteComment(c *gin.Context) {
	commentID := c.Param("id")
	err := co.c.DeleteComment(c.Request.Context(), commentID)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, nil)
}

func (co *commentAPI) getOnUser(c *gin.Context) {
	username := c.Param("username")
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	intpage, sErr := strconv.Atoi(page)
	if sErr != nil {
		sendMsg(c, http.StatusBadRequest, "page number should be a integer")
		return
	}
	cmts, err := co.c.GetCommentMadeOn(c.Request.Context(), username, model.UserComment, 10, int64(intpage))
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, cmts)
}

func (co *commentAPI) getOnBook(c *gin.Context) {
	id := c.Param("id")
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	intpage, sErr := strconv.Atoi(page)
	if sErr != nil {
		sendMsg(c, http.StatusBadRequest, "page number should be a integer")
		return
	}
	cmts, err := co.c.GetCommentMadeOn(c.Request.Context(), id, model.BookComment, 10, int64(intpage))
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, cmts)
}

func (co *commentAPI) getOnComment(c *gin.Context) {
	id := c.Param("id")
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	intpage, sErr := strconv.Atoi(page)
	if sErr != nil {
		sendMsg(c, http.StatusBadRequest, "page number should be a integer")
		return
	}
	cmts, err := co.c.GetCommentMadeOn(c.Request.Context(), id, model.CommentComment, 10, int64(intpage))
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, cmts)
}

func (co *commentAPI) update(c *gin.Context) {
	id := c.Param("id")
	var upcmt commentUpdate
	jErr := c.ShouldBindJSON(&upcmt)
	if jErr != nil {
		sendMsg(c, http.StatusBadRequest, "wrong json input")
		return
	}
	cmt, err := co.c.UpdateComment(c.Request.Context(), id, upcmt.Value)
	if err != nil {
		sendMsg(c, errors.ToHTTP[err.GetStatus()], err.Error())
		return
	}
	sendData(c, http.StatusOK, cmt)
}

type commentUpdate struct {
	Value string `json:"value,omitempty"`
}
