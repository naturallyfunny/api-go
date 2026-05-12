package api

import (
	"encoding/json"
	"io"
)

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"perPage,omitempty"`
	TotalItems int `json:"totalItems,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
}

// Response is the shared wire envelope for all API responses.
// Servers write it; clients decode into Response[T] for type-safe Data access.
// Use Response[json.RawMessage] to defer or skip Data decoding.
type Response[T any] struct {
	Success bool          `json:"success"`
	Code    Code          `json:"code"`
	Message string        `json:"message"`
	Data    T             `json:"data,omitempty"`
	Meta    *Meta         `json:"meta,omitempty"`
	Errors  []ErrorDetail `json:"errors,omitempty"`
}

// Decode reads a JSON response from r, unwraps the envelope, and returns the
// typed Data on success. On a non-success envelope it returns a *Error so
// callers can errors.As against it.
func Decode[T any](r io.Reader) (T, error) {
	var resp Response[T]
	if err := json.NewDecoder(r).Decode(&resp); err != nil {
		var zero T
		return zero, err
	}
	if !resp.Success {
		return resp.Data, &Error{
			Code:    resp.Code,
			Message: resp.Message,
			Details: resp.Errors,
		}
	}
	return resp.Data, nil
}
