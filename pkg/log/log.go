package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// Logger is a logger that support log level and structured logging
type Logger interface {

	// Debug is like fmt.Sprint()
	Debug(args ...interface{})
	// Debugf is like fmt.Sprintf()
	Debugf(format string, args ...interface{})

	// Info is like fmt.Sprint()
	Info(args ...interface{})
	// Infof is like fmt.Sprintf()
	Infof(format string, args ...interface{})

	Error(args ...interface{})
	// Error is like fmt.Sprint()
	Errorf(format string, args ...interface{})
	// Errorf is like fmt.Sprintf()
}

// New Creates a new logger using default configuration
func New() Logger {
	lgr := &logger{}
	return lgr
}

// NewForTest returns a new logger and the corresponding observed logs which
// can be used in unit to verify log entries
func NewForTest() (Logger, *observer.ObservedLogs) {
	core, recorder := observer.New(zapcore.InfoLevel)
	return &logger{zap.New(core).Sugar()}, recorder
}
