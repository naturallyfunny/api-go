package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"perPage,omitempty"`
	TotalItems int `json:"totalItems,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
}

type Response[T any] struct {
	Success bool          `json:"success"`
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Data    T             `json:"data,omitempty"`
	Meta    *Meta         `json:"meta,omitempty"`
	Errors  []ErrorDetail `json:"errors,omitempty"`
}

func writeJSON[T any](w http.ResponseWriter, status int, payload Response[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("api: failed to encode response: %v", err)
	}
}

func WriteSuccess[T any](w http.ResponseWriter, status int, code, message string, data T, meta *Meta) {
	writeJSON(w, status, Response[T]{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func writeFailure(w http.ResponseWriter, status int, code, message string, details []ErrorDetail) {
	writeJSON(w, status, Response[struct{}]{
		Success: false,
		Code:    code,
		Message: message,
		Errors:  details,
	})
}

func WriteError(w http.ResponseWriter, err error) {
	var apiErr *Error

	if errors.As(err, &apiErr) {
		if apiErr.Status >= http.StatusInternalServerError && apiErr.Err != nil {
			log.Printf("api: wrapped internal error: %v", apiErr.Err)
		}
		writeFailure(w, apiErr.Status, apiErr.Code, apiErr.Message, apiErr.Details)
		return
	}

	log.Printf("api: unhandled internal error: %v", err)
	writeFailure(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected error occurred", nil)
}
