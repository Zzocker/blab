package user

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/Zzocker/blab/core/ports"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/util"
)

type userCore struct {
	uStore ports.UserStorePort
}

func (u *userCore) Register(ctx context.Context, in Register) (*model.User, errors.E) {
	err := in.validate()
	if err != nil {
		return nil, err
	}
	usr := model.User{
		Username: in.Username,
		Details: model.UserDetails{
			Name:   in.Name,
			Age:    in.Age,
			Gender: in.Gender,
		},
		Contacts: []model.UserContact{
			{
				Type:  model.UserContactEmail,
				Value: in.EmailID,
			},
		},
		Rating:    model.UserRating{},
		CreatedOn: time.Now().Unix(),
		Password:  util.Hash(in.Password),
	}
	err = u.uStore.Store(ctx, usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (u *userCore) Get(ctx context.Context, username string) (*model.User, errors.E) {
	return u.uStore.Get(ctx, username)
}

func (u *userCore) Update(ctx context.Context, username string, reader io.Reader) (*model.User, errors.E) {
	usr, err := u.uStore.Get(ctx, username)
	if err != nil {
		return nil, err
	}
	jErr := json.NewDecoder(reader).Decode(usr)
	if jErr != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to decode")
	}
	usr.Username = username
	err = u.uStore.Update(ctx, usr.Username, *usr)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
