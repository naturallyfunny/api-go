package http

import (
	"encoding/json"
	"errors"
	"log"
	nethttp "net/http"

	"go.naturallyfunny.dev/api"
)

func writeJSON(w nethttp.ResponseWriter, status int, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Printf("api: failed to encode response: %v", err)
		nethttp.Error(w, `{"success":false,"code":13,"message":"Internal server error"}`, nethttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		log.Printf("api: failed to write response: %v", err)
	}
}

// WriteSuccess writes a successful JSON response envelope to w.
func WriteSuccess[T any](w nethttp.ResponseWriter, code api.Code, message string, data T, meta *api.Meta) {
	writeJSON(w, code.HTTPStatus(), api.Response[T]{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// WriteError writes an error JSON response envelope to w.
func WriteError(w nethttp.ResponseWriter, err error) {
	if err == nil {
		return
	}

	var apiErr *api.Error
	if errors.As(err, &apiErr) {
		if apiErr.Code.HTTPStatus() >= nethttp.StatusInternalServerError && apiErr.Err != nil {
			log.Printf("api: %s: %v", apiErr.Code, apiErr.Err)
		}
		writeJSON(w, apiErr.Code.HTTPStatus(), api.Response[json.RawMessage]{
			Success: false,
			Code:    apiErr.Code,
			Message: apiErr.Message,
			Errors:  apiErr.Details,
		})
		return
	}

	log.Printf("api: unhandled internal error: %v", err)
	writeJSON(w, nethttp.StatusInternalServerError, api.Response[json.RawMessage]{
		Success: false,
		Code:    api.Internal,
		Message: "An unexpected error occurred",
	})
}
