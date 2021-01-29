package model

type tokenType uint8

const (
	refreshToken tokenType = iota + 1
	accessToken
)

type Token struct {
	ID       string    `json:"id"`
	Type     tokenType `json:"type"`
	ExpireIn int64     `json:"expire_in"`
}
