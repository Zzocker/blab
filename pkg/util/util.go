package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(in string) string {
	return hex.EncodeToString(md5.New().Sum([]byte(in)))
}
