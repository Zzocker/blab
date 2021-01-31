package user

import (
	"context"
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
