package api

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func RequireHeader(name string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get(name) == "" {
				WriteError(w, NewError(http.StatusUnauthorized, "UNAUTHORIZED", "Missing required header: "+name))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
