package log

import (
	"github.com/sirupsen/logrus"
)

var (
	// L is global variable to be used by other pkg directly
	L Logger
)

type logger struct {
	*logrus.Logger
}
