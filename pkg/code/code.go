package code

import "net/http"

// code package defines all internal code used by this project

// Status is status code
type Status uint8

const (
	// CodeNotFound called when a item is not present
	CodeNotFound Status = iota + 1

	// CodeInternal called for internal server error
	CodeInternal

	// CodeAlreadyExists returned when trying to store a already exists document
	CodeAlreadyExists

	// CodeInvalidArgument returned when user send bad request
	CodeInvalidArgument

	// CodeUnauthorized :
	CodeUnauthorized
)

// ToHTTP convert internal status code into http status code
func ToHTTP(status Status) int {
	var hStatus int
	switch status {
	case CodeNotFound:
		hStatus = http.StatusNotFound
	case CodeInternal:
		hStatus = http.StatusInternalServerError
	case CodeAlreadyExists:
		hStatus = http.StatusConflict
	case CodeInvalidArgument:
		hStatus = http.StatusBadRequest
	case CodeUnauthorized:
		hStatus = http.StatusUnauthorized
	default:
		hStatus = http.StatusInternalServerError
	}
	return hStatus
}
