package comment

import (
	"context"

	"github.com/Zzocker/blab/core/ports"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/log"
)

type commentCore struct {
	cStore ports.CommentStore
}

func (c *commentCore) CommentOn(ctx context.Context, com CommentCreateInput) (*model.Comment, errors.E) {
	if err := com.validate(); err != nil {
		return nil, err
	}
	log.L.Info(ctx.Value("username"))
	return nil, errors.New(errors.CodeInternalErr, "implement me!!")
}
