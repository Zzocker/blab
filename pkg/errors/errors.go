package errors

// E represents error interface used throughout this application
type E interface {
	Error() string
	GetStatus() int
}

type err struct {
	code int
	msg  string
}

// New : Create a new error type
func New(code int, msg string) E {
	return &err{
		code: code,
		msg:  msg,
	}
}

func (e *err) Error() string {
	return e.msg
}
func (e *err) GetStatus() int {
	return e.code
}

const (
	CodeNotFound = iota + 1
	CodeAlreadyExists
	CodeInvalidArgument
	CodeInternalErr
)
