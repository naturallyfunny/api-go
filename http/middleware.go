package http

import (
	"context"
	nethttp "net/http"

	"go.naturallyfunny.dev/api"
)

// Middleware is an HTTP middleware function.
type Middleware func(nethttp.Handler) nethttp.Handler

// RequireHeader returns a Middleware that rejects requests missing the named header.
func RequireHeader(name string) Middleware {
	return func(next nethttp.Handler) nethttp.Handler {
		return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
			if r.Header.Get(name) == "" {
				WriteError(w, api.NewError(api.Unauthenticated, "Missing required header: "+name))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// HeaderToContext extracts a header value and uses the provided injector to
// update the request context. If the injector returns an error, it responds
// with failErr.
func HeaderToContext(
	headerName string,
	injector func(context.Context, string) (context.Context, error),
	failErr *api.Error,
) Middleware {
	return func(next nethttp.Handler) nethttp.Handler {
		return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
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
