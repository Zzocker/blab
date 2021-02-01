package book

import (
	"time"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/go-playground/validator"
)

var (
	va = validator.New()
)

type BookCreate struct {
	ISBN     string   `json:"isbn"`
	Author   string   `json:"author"`
	Genre    []string `json:"genre"`
	Username string   `json:"-"`
}

func (b BookCreate) validate() errors.E {
	var err errors.E
	if b.ISBN == "" {
		err = errors.New(errors.CodeInvalidArgument, "empty isbn")
	} else if b.Author == "" {
		err = errors.New(errors.CodeInvalidArgument, "empty author")
	} else if len(b.Genre) < 1 {
		err = errors.New(errors.CodeInvalidArgument, "should have attest one genre")
	}
	return err
}

func (b BookCreate) toBook() *model.Book {
	return &model.Book{
		ISBN: b.ISBN,
		Details: model.BookDetails{
			Author: b.Author,
			Genre:  b.Genre,
		},
		Ownership: model.BookOwnership{
			Owner:   b.Username,
			Current: b.Username,
		},
		Rating: model.Rating{
			Count: 0,
			Value: 0,
		},
		CreatedOn: time.Now().Unix(),
	}
}
