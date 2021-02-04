package errors

import "github.com/Zzocker/blab/pkg/code"

// E is error type
// Code to be sent to client
type E interface {
	GetCode() code.Status
	Error() string
}

// Err is for logging
// Err will be of top call stack
// <---- Datastore    -----> <Err from here> <Info Logging>
// <---- Code         ----->  <Info Logging>
// <---- HTTP Request -----> <Based on Code send correspond Message to client> <Info and Err Logging>
type e struct {
	Code code.Status
	Err  error
}

// InitErr : initate error, called by top most caller stack
func InitErr(err error, code code.Status) E {
	return &e{
		Code: code,
		Err:  err,
	}
}

func (e *e) GetCode() code.Status {
	return e.Code
}
func (e *e) Error() string {
	return e.Err.Error()
}
