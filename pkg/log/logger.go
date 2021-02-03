package log

import (
	"io"
	golog "log"
)

// Logger represents logger for this project
// Methods provided are
// Info , Error
type Logger interface {
	Info(f string, args ...interface{})
	Error(f string, args ...interface{})
}

// NewLogger creates a new log
// only to be used by this repo
func NewLogger(callerEnabled bool, infoWriter, errWriter io.Writer) Logger {
	if errWriter == nil {
		errWriter = infoWriter
	}
	var flag int
	if callerEnabled {
		flag = golog.Lshortfile | golog.LstdFlags
	}
	flag = flag | golog.Ltime | golog.Ldate | golog.Lmicroseconds
	return &logger{
		errLog:  golog.New(errWriter, errPrefixString, flag),
		infoLog: golog.New(errWriter, infoPrefixString, flag),
	}
}
