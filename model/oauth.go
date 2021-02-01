package model

import (
	"time"

	"github.com/google/uuid"
)

type tokenType uint8

const (
	RefreshToken tokenType = iota + 1
	AccessToken
)

var (
	refreshTokenExpiry = (30 * 24 * time.Hour).Seconds()
	accessTokenExpiry  = (7 * 24 * time.Hour).Seconds()
)

type Token struct {
	ID        string    `json:"id" bson:"id"`
	Username  string    `json:"username" bson:"username"`
	Type      tokenType `json:"type" bson:"type"`
	ExpireIn  int64     `json:"expire_in" bson:"expire_in"`
	CreatedAt int64     `json:"created_at" bson:"created_at"`
}

func NewRefreshToken(username string) Token {
	return Token{
		ID:        uuid.New().String(),
		Username:  username,
		Type:      RefreshToken,
		ExpireIn:  int64(refreshTokenExpiry),
		CreatedAt: time.Now().Unix(),
	}
}

func NewAccessToken(username string) Token {
	return Token{
		ID:        uuid.New().String(),
		Username:  username,
		Type:      AccessToken,
		ExpireIn:  int64(accessTokenExpiry),
		CreatedAt: time.Now().Unix(),
	}
}
