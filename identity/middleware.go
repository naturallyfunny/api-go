package identity

import (
	"net/http"

	"go.ibnfadl.com/api"
)

func RequireUserID(next http.Handler) http.Handler {
	return RequireUserIDWithHeader("x-user-id")(next)
}

func RequireUserIDWithHeader(headerName string) api.Middleware {
	return func(next http.Handler) http.Handler {
		return api.RequireHeader(headerName)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := r.Header.Get(headerName)

			ctx, err := api.NewContextWithUserID(r.Context(), userID)
			if err != nil {
				api.WriteError(w, api.NewError(http.StatusUnauthorized, "UNAUTHORIZED", "Invalid user identity"))
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}))
	}
}
