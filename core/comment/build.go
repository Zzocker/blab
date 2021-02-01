package comment

import (
	"github.com/Zzocker/blab/adapters"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
)

func NewCommentCore(conf config.C) (*commentCore, errors.E) {
	store, err := adapters.CreateCommentStore(conf.Core.Book.BookStoreConf)
	if err != nil {
		return nil, err
	}
	return &commentCore{cStore: store}, nil
}
