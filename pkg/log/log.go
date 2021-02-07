package log

import "io"

type Logger interface {
	// [requestID] [packageName] [msg]
	Debugf(reqID int64, pkg string, f string, args ...interface{})
	Infof(reqID int64, pkg string, f string, args ...interface{})
	Errorf(reqID int64, pkg string, f string, args ...interface{})
}

// New create a new logger
func New(writer io.Writer, lvl Level) Logger {
	return newBasicLogger(writer, lvl)
}
