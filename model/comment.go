package model

type commentType string

const (
	// userComment is comment on user
	userComment commentType = "user_comment"

	// bookComment is comment on book
	bookComment commentType = "book_comment"

	// userComment is comment on user
	commentComment commentType = "comment_comment"
)

type Comment struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	On    string `json:"on"`    // id of which comment is being made
	By    string `json:"by"`    // username of user who made this comment
	Value string `json:"value"` // actual content of comment
}
