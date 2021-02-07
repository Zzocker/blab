package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Level log level
type Level int

const (
	// Debug = -1 means by default debug is disabled
	Debug Level = iota - 1

	// Info is default log level
	Info

	// Error will always be there
	Error
)
const (
	skipCallerDepth = 2
)
const (
	debugPrefix = "[ DEBUG ] "
	infoPrefix  = "[ INFO  ] "
	errorPrefix = "[ ERROR ] "

	debugColor = "\033[0;36m%s\033[0m"
	infoColor  = "\033[1;34m%s\033[0m"
	errorColor = "\033[1;31m%s\033[0m"
)

type logger struct {
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
	lvl         Level
}

func newBasicLogger(writer io.Writer, lvl Level) *logger {
	if writer == nil {
		writer = os.Stdout
	}

	return &logger{
		debugLogger: log.New(writer, fmt.Sprintf(debugColor, debugPrefix), 0),
		errorLogger: log.New(writer, fmt.Sprintf(errorColor, errorPrefix), 0),
		infoLogger:  log.New(writer, fmt.Sprintf(infoColor, infoPrefix), 0),
		lvl:         lvl,
	}
}

func (l *logger) Debugf(reqID int, pkg string, f string, args ...interface{}) {
	l.output(Debug, []byte(fmt.Sprintf("[ %d ] [ %s ] %s", reqID, pkg, fmt.Sprintf(f, args...))))
}
func (l *logger) Infof(reqID int, pkg string, f string, args ...interface{}) {
	l.output(Info, []byte(fmt.Sprintf("[ %d ] [ %s ] %s", reqID, pkg, fmt.Sprintf(f, args...))))
}
func (l *logger) Errorf(reqID int, pkg string, f string, args ...interface{}) {
	l.output(Error, []byte(fmt.Sprintf("[ %d ] [ %s ] %s", reqID, pkg, fmt.Sprintf(f, args...))))
}

func (l *logger) output(lvl Level, msg []byte) {
	if lvl < l.lvl {
		return
	}
	switch lvl {
	case Debug:
		l.debugLogger.Output(skipCallerDepth, string(msg))
	case Info:
		l.infoLogger.Output(skipCallerDepth, string(msg))
	case Error:
		l.errorLogger.Output(skipCallerDepth, string(msg))
	}
}
