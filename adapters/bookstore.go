package adapters

import (
	"context"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/datastore"
	"github.com/Zzocker/blab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	primaryKeyName = "isbn"
)

type bookStore struct {
	db datastore.SmartDS
}

func (b *bookStore) Store(ctx context.Context, book model.Book) errors.E {
	return b.db.Store(ctx, book)
}
func (b *bookStore) Get(ctx context.Context, isbn string) (*model.Book, errors.E) {
	raw, err := b.db.Get(ctx, primaryKeyName, isbn)
	if err != nil {
		return nil, err
	}
	var book model.Book
	bErr := bson.Unmarshal(raw, &book)
	if bErr != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to unmarshal raw data")
	}
	return &book, nil
}
func (b *bookStore) Update(ctx context.Context, isbn string, book model.Book) errors.E {
	return b.db.Update(ctx, primaryKeyName, isbn, book)
}
func (b *bookStore) Delete(ctx context.Context, isbn string) errors.E {
	return b.db.Delete(ctx, primaryKeyName, isbn)
}
func (b *bookStore) Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber, perPage int64) ([]model.Book, errors.E) {
	raws, err := b.db.Query(ctx, sortKey, query, pageNumber, perPage)
	if err != nil {
		return nil, err
	}
	books := make([]model.Book, len(raws))
	for i := range raws {
		bson.Unmarshal(raws[i], &books[i])
	}
	return books, nil
}
