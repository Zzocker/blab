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
	Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber, perPage int64) ([]model.Book, errors.E)
}

// UserStorePort represents userprofile database, which will be implemented by user store adapter
// and used by BookCore
type UserStorePort interface {
	Store(ctx context.Context, user model.User) errors.E
	Get(ctx context.Context, username string) (*model.User, errors.E)
	Update(ctx context.Context, username string, user model.User) errors.E
	Delete(ctx context.Context, username string) errors.E
	Query(ctx context.Context, query map[string]interface{}, pageNumber, perPage int) ([]model.User, errors.E)
}

// OAuthStore represents oauth database, which will be implemented by oauth store adapter
// and used by oauth core
type OAuthStore interface {
	Store(ctx context.Context, token model.Token) errors.E
	Get(ctx context.Context, tokenID string) (*model.Token, errors.E)
	Delete(ctx context.Context, tokeID string) errors.E
}

// CommentStore represents comment database, which will be implemented by comment store adapter
// and used by commentCore
type CommentStore interface {
	Store(ctx context.Context, comment model.Comment) errors.E
	Get(ctx context.Context, commentID string) (*model.Comment, errors.E)
	Update(ctx context.Context, commentID string, comment model.Comment) errors.E
	Delete(ctx context.Context, comment string) errors.E
	Query(ctx context.Context, query map[string]interface{}, pageNumber, perPage int) ([]model.Comment, errors.E)
}
