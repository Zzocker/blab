package user

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/Zzocker/blab/pkg/errors"
	"github.com/go-playground/validator"
)

var validate = validator.New()

type gender string

const (
	genderMale   gender = "male"
	genderFemale        = "female"
)

type User struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Name     string `json:"name,omitempty" bson:name`
	Gender   gender `json:"gender,omitempty" bson:"gender"`
	EmailID  string `json:"email_id,omitempty" bson:"email_id"`
	Password string `json:"-" bson:"password"`
}

// ValidateRegister will validate user when regsitering a user
// will also hash the password
func (u *User) ValidateRegister() errors.E {
	var err errors.E
	if u.Username == "" {
		err = errors.New(http.StatusBadRequest, "empty username")
	} else if u.Name == "" {
		err = errors.New(http.StatusBadRequest, "empty name")
	} else if inErr := validate.Var(u.EmailID, "required,email"); inErr != nil {
		err = errors.New(http.StatusBadRequest, "invalid email")
	} else if u.Password == "" {
		err = errors.New(http.StatusBadRequest, "empty password")
	}
	switch u.Gender {
	case genderMale:
	case genderFemale:
	default:
		err = errors.New(http.StatusBadRequest, "invalid gender type")
	}
	u.Password = hex.EncodeToString(md5.New().Sum([]byte(u.Password)))
	if err != nil {
		return err
	}
	return nil
}
