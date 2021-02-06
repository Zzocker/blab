package core

import (
	"crypto/md5"
	"encoding/hex"

	"gopkg.in/go-playground/validator.v9"
)

var (
	validate = validator.New()
)

func hash(s string) string {
	return hex.EncodeToString(md5.New().Sum([]byte(s)))
}
