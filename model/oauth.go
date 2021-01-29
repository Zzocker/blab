package model

type tokenType uint8

const (
	refreshToken tokenType = iota + 1
	accessToken
)

type Token struct {
	ID       string    `json:"id" bson:"id"`
	Type     tokenType `json:"type" bson:"type"`
	ExpireIn int64     `json:"expire_in" bson:"expire_in"`
}
