package api

import (
	"fmt"
)

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type Error struct {
	Status  int
	Code    string
	Message string
	Details []ErrorDetail
	Err     error
}

func NewError(status int, code, message string) *Error {
	return &Error{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func (e *Error) WithDetails(details []ErrorDetail) *Error {
	e.Details = details
	return e
}

func (e *Error) WithError(err error) *Error {
	e.Err = err
	return e
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("api: %s (%d): %s - %v", e.Code, e.Status, e.Message, e.Err)
	}
	return fmt.Sprintf("api: %s (%d): %s", e.Code, e.Status, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}
