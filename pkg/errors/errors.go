package errors

// E represents error interface used throughout this application
type E interface {
	Error() string
	GetStatus() codes
}

type err struct {
	code codes
	msg  string
}

// New : Create a new error type
func New(code codes, msg string) E {
	return &err{
		code: code,
		msg:  msg,
	}
}

func (e *err) Error() string {
	return e.msg
}
func (e *err) GetStatus() codes {
	return e.code
}

type codes uint8

const (
	CodeNotFound codes = iota + 1
	CodeAlreadyExists
	CodeInvalidArgument
	CodeInternalErr
)
