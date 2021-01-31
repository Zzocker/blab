package errors

import "net/http"

// E represents error interface used throughout this application
type E interface {
	Error() string
	GetStatus() code
}

type err struct {
	code code
	msg  string
}

// New : Create a new error type
func New(code code, msg string) E {
	return &err{
		code: code,
		msg:  msg,
	}
}

func (e *err) Error() string {
	return e.msg
}
func (e *err) GetStatus() code {
	return e.code
}

type code uint8

const (
	CodeNotFound code = iota + 1
	CodeAlreadyExists
	CodeInvalidArgument
	CodeInternalErr
	CodeUnauthorized
)

var (
	ToHTTP = map[code]int{
		CodeNotFound:        http.StatusNotFound,
		CodeAlreadyExists:   http.StatusConflict,
		CodeInvalidArgument: http.StatusBadRequest,
		CodeInternalErr:     http.StatusInternalServerError,
		CodeUnauthorized:    http.StatusUnauthorized,
	}
)
