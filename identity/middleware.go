package identity

import (
	"net/http"

	"go.naturallyfunny.dev/api"
)

func WithUserID(next http.Handler) http.Handler {
	return WithUserIDFromHeader("x-user-id")(next)
}

func WithUserIDFromHeader(headerName string) api.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := r.Header.Get(headerName)
			ctx, err := NewContextWithUserID(r.Context(), userID)
			if err != nil {
				api.WriteError(w, api.NewError(api.Unauthenticated, "Missing required identity header"))
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WithSessionID(next http.Handler) http.Handler {
	return WithSessionIDFromHeader("x-session-id")(next)
}

func WithSessionIDFromHeader(headerName string) api.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionID := r.Header.Get(headerName)
			ctx, err := NewContextWithSessionID(r.Context(), sessionID)
			if err != nil {
				api.WriteError(w, api.NewError(api.Unauthenticated, "Missing required session identity header"))
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WithUserTimezone(next http.Handler) http.Handler {
	return WithUserTimezoneFromHeader("x-user-timezone")(next)
}

func WithUserTimezoneFromHeader(headerName string) api.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tz := r.Header.Get(headerName)
			ctx, err := NewContextWithTimezone(r.Context(), tz)
			if err != nil {
				api.WriteError(w, api.NewError(api.InvalidArgument, "Missing or invalid timezone header: expected a valid IANA timezone name (e.g. Asia/Jakarta)"))
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
