package book

import (
	"context"
	"encoding/json"
	"io"

	"github.com/Zzocker/blab/core/ports"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

type bookCore struct {
	bStore ports.BookStorePort
}

func (b *bookCore) AddBook(ctx context.Context, in BookCreate) (*model.Book, errors.E) {
	err := in.validate()
	if err != nil {
		return nil, err
	}
	bk := in.toBook()
	err = b.bStore.Store(ctx, *bk)
	if err != nil {
		return nil, err
	}
	return bk, err
}
func (b *bookCore) Get(ctx context.Context, isbn string) (*model.Book, errors.E) {
	return b.bStore.Get(ctx, isbn)
}

// TODO : check if book is owned by user making this request
// Use context to get username
func (b *bookCore) Update(ctx context.Context, isbn string, reader io.Reader) (*model.Book, errors.E) {
	bk, err := b.bStore.Get(ctx, isbn)
	if err != nil {
		return nil, err
	}
	jErr := json.NewDecoder(reader).Decode(bk)
	if jErr != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to decode json")
	}
	err = b.bStore.Update(ctx, isbn, *bk)
	if err != nil {
		return nil, err
	}
	return bk, nil
}

// TODO : check if book is owned by user making this request
func (b *bookCore) Remove(ctx context.Context, isbn string) errors.E {
	return b.bStore.Delete(ctx, isbn)
}
