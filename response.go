package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"perPage,omitempty"`
	TotalItems int `json:"totalItems,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
}

// Response is the shared wire envelope for all Avagenc API responses.
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

func writeJSON(w http.ResponseWriter, status int, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Printf("api: failed to encode response: %v", err)
		http.Error(w, `{"success":false,"code":13,"message":"Internal server error"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		log.Printf("api: failed to write response: %v", err)
	}
}

func WriteSuccess[T any](w http.ResponseWriter, code Code, message string, data T, meta *Meta) {
	writeJSON(w, code.HTTPStatus(), Response[T]{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func WriteError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	var apiErr *Error
	if errors.As(err, &apiErr) {
		if apiErr.Code.HTTPStatus() >= http.StatusInternalServerError && apiErr.Err != nil {
			log.Printf("api: %s: %v", apiErr.Code, apiErr.Err)
		}
		writeJSON(w, apiErr.Code.HTTPStatus(), Response[json.RawMessage]{
			Success: false,
			Code:    apiErr.Code,
			Message: apiErr.Message,
			Errors:  apiErr.Details,
		})
		return
	}

	log.Printf("api: unhandled internal error: %v", err)
	writeJSON(w, http.StatusInternalServerError, Response[json.RawMessage]{
		Success: false,
		Code:    Internal,
		Message: "An unexpected error occurred",
	})
}
