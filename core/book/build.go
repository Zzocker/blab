package book

import (
	"github.com/Zzocker/blab/adapters"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
)

func NewBookCore(conf config.C) (*bookCore, errors.E) {
	store, err := adapters.CreateBookStore(conf.Core.Book.BookStoreConf)
	if err != nil {
		return nil, err
	}
	return &bookCore{bStore: store}, nil
}
