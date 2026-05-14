package http

import (
	"encoding/json"
	"log"
	nethttp "net/http"
)

// WriteJSON writes successful raw payloads directly using the application/json media type.
func WriteJSON(w nethttp.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("api: failed to encode response: %v", err)
		}
	}
}

// WriteProblem writes raw error payloads directly using the application/json media type.
func WriteProblem(w nethttp.ResponseWriter, status int, problem any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if problem != nil {
		if err := json.NewEncoder(w).Encode(problem); err != nil {
			log.Printf("api: failed to encode problem: %v", err)
		}
	}
}
