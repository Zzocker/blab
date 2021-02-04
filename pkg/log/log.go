package log

import (
	"fmt"
	"log"
)

type level uint8

const (
	errorLevel level = iota + 1
	infoLevel

	// caller depth
	callerDepth = 3
)

var (
	errPrefixString  = "[ ERROR ] "
	infoPrefixString = "[ INFO  ] "
)

type logger struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func (l *logger) Info(f string, args ...interface{}) {
	l.output(infoLevel, fmt.Sprintf(f, args...))
}
func (l *logger) Error(f string, args ...interface{}) {
	l.output(errorLevel, fmt.Sprintf(f, args...))
}

func (l *logger) output(lvl level, msg string) {
	switch lvl {
	case errorLevel:
		l.errLog.Output(callerDepth, msg)
	case infoLevel:
		l.infoLog.Output(callerDepth, msg)
	}
}
