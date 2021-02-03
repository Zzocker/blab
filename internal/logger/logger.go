package logger

import (
	"os"

	"github.com/Zzocker/blab/pkg/log"
)

// L is project level logger
// to de used by every pkg
var L log.Logger

func init() {
	L = log.NewLogger(true, os.Stdout, nil)
}
