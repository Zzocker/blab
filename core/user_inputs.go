package core

import (
	"fmt"
	"time"

	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/code"
	"github.com/Zzocker/blab/pkg/errors"
)

type RegisterUserInput struct {
	Username string           `json:"username"`
	Name     string           `json:"name"`
	DOB      int64            `json:"dob"`
	Gender   model.GenderType `json:"gender"`
	Email    string           `json:"email"`
}

func (r RegisterUserInput) validate(password string) errors.E {
	logger.L.Info(userloggerPrefix, "validating register input")
	var err errors.E
	if r.Username == "" {
		err = errors.InitErr(fmt.Errorf("empty username"), code.CodeInvalidArgument)
	} else if r.Name == "" {
		err = errors.InitErr(fmt.Errorf("empty name"), code.CodeInvalidArgument)
	} else if r.DOB <= 0 {
		err = errors.InitErr(fmt.Errorf("dob should be epoch value"), code.CodeInvalidArgument)
	} else if password == "" {
		err = errors.InitErr(fmt.Errorf("empty password"), code.CodeInvalidArgument)
	} else if r.Email == "" {
		err = errors.InitErr(fmt.Errorf("empty emailID"), code.CodeInvalidArgument)
	} else if vErr := validate.Var(r.Email, "email"); vErr != nil {
		err = errors.InitErr(fmt.Errorf("invalid emailID"), code.CodeInvalidArgument)
	}

	if r.Gender == model.GenderTypeMale {
	} else if r.Gender == model.GenderTypeFemale {
	} else {
		err = errors.InitErr(fmt.Errorf("invalid gender type"), code.CodeInvalidArgument)
	}
	if err != nil {
		return err
	}
	logger.L.Info(userloggerPrefix, "valide register input")
	return nil
}

func (r RegisterUserInput) toUser(hashPass string) model.User {
	return model.User{
		Username: r.Username,
		Details: model.UserDetails{
			Name:   r.Name,
			DOB:    r.DOB,
			Gender: r.Gender,
		},
		Contacts: []model.UserContact{
			{
				Type:  model.UserContactTypeEmail,
				Value: r.Email,
			},
		},
		CreatedOn: time.Now().Unix(),
		Password:  hashPass,
	}
}
