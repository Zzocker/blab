package log

import (
	"github.com/sirupsen/logrus"
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

	// Warn is like fmt.Sprint()
	Warn(args ...interface{})
	// Warnf is like fmt.Sprintf()
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	// Error is like fmt.Sprint()
	Errorf(format string, args ...interface{})
	// Errorf is like fmt.Sprintf()
}

// New Creates a new logger using default configuration
func New() Logger {
	l := logrus.New()
	l.WriterLevel(logrus.DebugLevel)
	return &logger{l}
}

// WithFields :
func WithFields(l Logger, fields map[string]interface{}) Logger {
	return l.(*logger).WithFields(fields)
}