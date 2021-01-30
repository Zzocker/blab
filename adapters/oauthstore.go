package adapters

import (
	"context"
	"encoding/json"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/datastore"
	"github.com/Zzocker/blab/pkg/errors"
)

type oauthStore struct {
	db datastore.DumbDS
}

func (o *oauthStore) Store(ctx context.Context, token model.Token) errors.E {
	raw, err := json.Marshal(token)
	if err != nil {
		return errors.New(errors.CodeInternalErr, "failed to marshal token")
	}
	return o.db.Store(ctx, token.ID, raw, token.ExpireIn)
}
func (o *oauthStore) Get(ctx context.Context, tokenID string) (*model.Token, errors.E) {
	raw, err := o.db.Get(ctx, tokenID)
	if err != nil {
		return nil, err
	}
	var token model.Token
	jErr := json.Unmarshal(raw, &token)
	if jErr != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to unmarshal token")
	}
	return &token, nil
}
func (o *oauthStore) Delete(ctx context.Context, tokenID string) errors.E {
	return o.db.Delete(ctx, tokenID)
}
