package adapters

import (
	"context"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/datastore"
	"github.com/Zzocker/blab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	commentPrimaryKeyName = "id"
)

type commentStore struct {
	db datastore.SmartDS
}

func (b *commentStore) Store(ctx context.Context, comment model.Comment) errors.E {
	return b.db.Store(ctx, comment)
}
func (b *commentStore) Get(ctx context.Context, id string) (*model.Comment, errors.E) {
	raw, err := b.db.Get(ctx, commentPrimaryKeyName, id)
	if err != nil {
		return nil, err
	}
	var comment model.Comment
	bErr := bson.Unmarshal(raw, &comment)
	if bErr != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to unmarshal raw data")
	}
	return &comment, nil
}
func (b *commentStore) Update(ctx context.Context, id string, comment model.Comment) errors.E {
	return b.db.Update(ctx, commentPrimaryKeyName, id, comment)
}
func (b *commentStore) Delete(ctx context.Context, id string) errors.E {
	return b.db.Delete(ctx, commentPrimaryKeyName, id)
}
func (b *commentStore) Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber, perPage int64) ([]model.Comment, errors.E) {
	raws, err := b.db.Query(ctx, sortKey, query, pageNumber, perPage)
	if err != nil {
		return nil, err
	}
	comments := make([]model.Comment, len(raws))
	for i := range raws {
		bson.Unmarshal(raws[i], &comments[i])
	}
	return comments, nil
}
