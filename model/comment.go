package model

type CommentType string

const (
	// userComment is comment on user
	UserComment CommentType = "user_comment"

	// bookComment is comment on book
	BookComment CommentType = "book_comment"

	// userComment is comment on user
	CommentComment CommentType = "comment_comment"
)

type Comment struct {
	ID    string      `json:"id" bson:"id"`
	Type  CommentType `json:"type" bson:"type"`
	On    string      `json:"on" bson:"on"`       // id on which comment is being made
	By    string      `json:"by" bson:"by"`       // username of user who made this comment
	Value string      `json:"value" bson:"value"` // actual content of comment
	When  int64       `json:"when" bson:"when"`
}
