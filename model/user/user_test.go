package user

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRegister(t *testing.T) {
	u := User{
		Username: "username",
		Name:     "name",
		Gender:   genderMale,
		EmailID:  "pkspritam@gmail.com",
		Password: "password",
	}
	err := u.ValidateRegister()
	assert.NoError(t, err)
	assert.NotEqual(t, "password", u.Password)

	u.Username = ""
	err = u.ValidateRegister()
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, err.GetStatus())

	u.Name = ""
	err = u.ValidateRegister()
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, err.GetStatus())

	u.Gender = ""
	err = u.ValidateRegister()
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, err.GetStatus())

	u.EmailID = "kdsjhkfksdjf.com"
	err = u.ValidateRegister()
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, err.GetStatus())

}
