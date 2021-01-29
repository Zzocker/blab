package core

import (
	"context"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

// BookStorePort represent book store port, which will be implemented by book store adapter
// and used by BookCore
type BookStorePort interface {
	Store(ctx context.Context, book model.Book) errors.E
	Get(ctx context.Context, isbn string) (*model.Book, errors.E)
	Update(ctx context.Context, isbn string, book model.Book) errors.E
	Delete(ctx context.Context, isbn string) errors.E
	Query(ctx context.Context, query map[string]interface{}, pageNumber, perPage int) ([]model.Book, errors.E)
}

// UserStorePort represents userprofile database, which will be implemented by user store adapter
// and used by BookCore
type UserStorePort interface {
	Store(ctx context.Context,user )
}
