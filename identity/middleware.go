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
