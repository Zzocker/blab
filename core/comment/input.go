package comment

import (
	"time"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/google/uuid"
)

type CommentCreateInput struct {
	Comment string `json:"comment"`
	On      string `json:"on"`
}

func (c CommentCreateInput) validate() errors.E {
	var err errors.E
	if c.Comment == "" {
		err = errors.New(errors.CodeInvalidArgument, "empty comment")
	} else if c.Comment == "" {
		err = errors.New(errors.CodeInvalidArgument, "empty comment on")
	}
	return err
}

func (c CommentCreateInput) toComment(commentType model.CommentType, username string) *model.Comment {
	return &model.Comment{
		ID:    uuid.New().String(),
		Type:  commentType,
		On:    c.On,
		By:    username,
		Value: c.Comment,
		When:  time.Now().Unix(),
	}
}
