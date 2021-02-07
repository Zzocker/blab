package logger

import (
	"os"

	"github.com/Zzocker/blab/pkg/log"
)

// L is application level logger used by all package
var L log.Logger

// Register will initiate L
func Register(level log.Level) {
	L = log.New(os.Stdout, level)
}
