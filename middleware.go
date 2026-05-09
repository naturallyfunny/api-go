package api

import (
	"context"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func RequireHeader(name string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get(name) == "" {
				WriteError(w, NewError(Unauthenticated, "Missing required header: "+name))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// HeaderToContext extracts a header value and uses the provided injector function
// to update the request context. If the injector returns an error, it responds
// with the provided failErr.
func HeaderToContext(
	headerName string,
	injector func(context.Context, string) (context.Context, error),
	failErr *Error,
) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Header.Get(headerName)
			ctx, err := injector(r.Context(), val)
			if err != nil {
				WriteError(w, failErr)
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
