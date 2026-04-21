package api

import (
	"fmt"
	"net/http"
)

// Code is a gRPC-style status code.
type Code uint32

const (
	OK                 Code = 0
	Cancelled          Code = 1
	Unknown            Code = 2
	InvalidArgument    Code = 3
	DeadlineExceeded   Code = 4
	NotFound           Code = 5
	AlreadyExists      Code = 6
	PermissionDenied   Code = 7
	ResourceExhausted  Code = 8
	FailedPrecondition Code = 9
	Aborted            Code = 10
	OutOfRange         Code = 11
	Unimplemented      Code = 12
	Internal           Code = 13
	Unavailable        Code = 14
	DataLoss           Code = 15
	Unauthenticated    Code = 16
)

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Cancelled:
		return "CANCELLED"
	case Unknown:
		return "UNKNOWN"
	case InvalidArgument:
		return "INVALID_ARGUMENT"
	case DeadlineExceeded:
		return "DEADLINE_EXCEEDED"
	case NotFound:
		return "NOT_FOUND"
	case AlreadyExists:
		return "ALREADY_EXISTS"
	case PermissionDenied:
		return "PERMISSION_DENIED"
	case ResourceExhausted:
		return "RESOURCE_EXHAUSTED"
	case FailedPrecondition:
		return "FAILED_PRECONDITION"
	case Aborted:
		return "ABORTED"
	case OutOfRange:
		return "OUT_OF_RANGE"
	case Unimplemented:
		return "UNIMPLEMENTED"
	case Internal:
		return "INTERNAL"
	case Unavailable:
		return "UNAVAILABLE"
	case DataLoss:
		return "DATA_LOSS"
	case Unauthenticated:
		return "UNAUTHENTICATED"
	default:
		return fmt.Sprintf("CODE(%d)", uint32(c))
	}
}

func (c Code) HTTPStatus() int {
	switch c {
	case OK:
		return http.StatusOK
	case Cancelled:
		return 499
	case Unknown:
		return http.StatusInternalServerError
	case InvalidArgument:
		return http.StatusBadRequest
	case DeadlineExceeded:
		return http.StatusGatewayTimeout
	case NotFound:
		return http.StatusNotFound
	case AlreadyExists:
		return http.StatusConflict
	case PermissionDenied:
		return http.StatusForbidden
	case ResourceExhausted:
		return http.StatusTooManyRequests
	case FailedPrecondition:
		return http.StatusBadRequest
	case Aborted:
		return http.StatusConflict
	case OutOfRange:
		return http.StatusBadRequest
	case Unimplemented:
		return http.StatusNotImplemented
	case Internal:
		return http.StatusInternalServerError
	case Unavailable:
		return http.StatusServiceUnavailable
	case DataLoss:
		return http.StatusInternalServerError
	case Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type Error struct {
	Code    Code
	Message string
	Details []ErrorDetail
	Err     error
}

func NewError(code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) WithDetails(details []ErrorDetail) *Error {
	e2 := *e
	e2.Details = details
	return &e2
}

func (e *Error) WithError(err error) *Error {
	e2 := *e
	e2.Err = err
	return &e2
}

func (e *Error) Error() string {
	if e == nil {
		return "api: <nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("api: %s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("api: %s: %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}
