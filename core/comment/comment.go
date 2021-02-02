package comment

import (
	"context"

	"github.com/Zzocker/blab/core/ports"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

type commentCore struct {
	cStore ports.CommentStore
}

// Before commenting check if user,book or comment exists TODO
func (c *commentCore) CommentOn(ctx context.Context, com CommentCreateInput, comType model.CommentType) (*model.Comment, errors.E) {
	if err := com.validate(); err != nil {
		return nil, err
	}
	cmt := com.toComment(comType, ctx.Value("username").(string))
	err := c.cStore.Store(ctx, *cmt)
	if err != nil {
		return nil, err
	}
	return cmt, nil
}

func (c *commentCore) GetComment(ctx context.Context, commentID string) (*model.Comment, errors.E) {
	return c.cStore.Get(ctx, commentID)
}

func (c *commentCore) DeleteComment(ctx context.Context, commentID string) errors.E {
	return c.cStore.Delete(ctx, commentID)
}

func (c *commentCore) GetCommentMadeOn(ctx context.Context, onID string, comType model.CommentType, perPage, pagNumber int64) ([]model.Comment, errors.E) {
	query := map[string]interface{}{
		"on":   onID,
		"type": comType,
	}
	// TODO give this perPage and pageCount
	return c.cStore.Query(ctx, "when", query, pagNumber, perPage)
}

func (c *commentCore) UpdateComment(ctx context.Context, cmtID, updateString string) (*model.Comment, errors.E) {
	cmt, err := c.cStore.Get(ctx, cmtID)
	if err != nil {
		return nil, err
	}
	cmt.Value = updateString
	err = c.cStore.Update(ctx, cmtID, *cmt)
	if err != nil {
		return nil, err
	}
	return cmt, nil
}
