package user

import (
	"context"

	"github.com/Zzocker/blab/pkg/log"

	"net/http"

	"github.com/Zzocker/blab/model/user"
	"github.com/Zzocker/blab/pkg/response"
	"github.com/gin-gonic/gin"
)

// RegisterHandlers will register endpoints for user api
func RegisterHandlers(r *gin.Engine, userAccess *user.DSAccess, l log.Logger) {
	intRouter := &router{userAccess, l}
	r.POST("/user/register", intRouter.register)
	r.GET("/user/:username", intRouter.get)
	r.PUT("/user", intRouter.put)
	r.PATCH("/user/:username", intRouter.patch)
	r.DELETE("/user/:username", intRouter.delete)
}

type router struct {
	*user.DSAccess
	l log.Logger
}

func (r *router) register(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	var usr user.User
	if err := c.ShouldBindJSON(&usr); err != nil {
		r.l.Error(err)
		res.SetCode(http.StatusBadRequest, "improper json request")
		return
	}
	usr.Password = c.GetHeader("secret")
	if err := usr.ValidateRegister(); err != nil {
		r.l.Error(err)
		res.SetCode(err.GetStatus(), err.Error())
		return
	}
	if err := r.Create(context.Background(), usr); err != nil {
		r.l.Error(err)
		res.SetCode(err.GetStatus(), err.Error())
		return
	}
	res.SetData(http.StatusCreated, nil)
}

func (r *router) get(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusServiceUnavailable, "StatusServiceUnavailable")
	u, err := r.Get(context.Background(), c.Param("username"))
	if err != nil {
		r.l.Error(err)
		res.SetCode(err.GetStatus(), err.Error())
		return
	}
	res.SetData(http.StatusOK, u)
}

func (r *router) put(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	res.SetCode(http.StatusServiceUnavailable, "StatusServiceUnavailable")
}

func (r *router) patch(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	oldUser, err := r.Get(context.Background(), c.Param("username"))
	if err != nil {
		r.l.Error(err)
		res.SetCode(err.GetStatus(), err.Error())
		return
	}
	if err := c.BindJSON(oldUser); err != nil {
		r.l.Error(err)
		res.SetCode(http.StatusInternalServerError, "bson error")
		return
	}
	if err := r.Update(context.TODO(), oldUser); err != nil {
		r.l.Error(err)
		res.SetCode(err.GetStatus(), err.Error())
		return
	}
	res.SetData(http.StatusOK, oldUser)
}

func (r *router) delete(c *gin.Context) {
	res := response.New()
	defer res.Send(c)
	if err := r.Delete(context.Background(), c.Param("username")); err != nil {
		r.l.Error(err)
		res.SetCode(err.GetStatus(), err.Error())
		return
	}
	res.SetData(http.StatusOK, nil)
}
