package code

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
)
