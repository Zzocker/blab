package user

import (
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/go-playground/validator"
)

var (
	v = validator.New()
)

type Register struct {
	Username string `json:"username"`
	EmailID  string `json:"email_id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Password string `json:"-"`
}

func (r Register) validate() (err errors.E) {
	if r.Username == "" {
		err = errors.New(errors.CodeInvalidArgument, "empty username")
	} else if vErr := v.Var(r.EmailID, "required,email"); vErr != nil {
		err = errors.New(errors.CodeInvalidArgument, "invalid email address")
	} else if r.Name == "" {
		err = errors.New(errors.CodeInvalidArgument, "empty name")
	} else if r.Age <= 0 {
		err = errors.New(errors.CodeInvalidArgument, "invalid age")
	} else if r.Password == "" {
		err = errors.New(errors.CodeInvalidArgument, "weak password")
	}

	if r.Gender == "male" {
	} else if r.Gender == "female" {
	} else {
		err = errors.New(errors.CodeInvalidArgument, "invalid gender type")
	}
	return
}
