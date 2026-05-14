package http

import (
	"context"
	nethttp "net/http"
)

type Middleware func(nethttp.Handler) nethttp.Handler

func RequireHeader(name string) Middleware {
	return func(next nethttp.Handler) nethttp.Handler {
		return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
			if r.Header.Get(name) == "" {
				WriteProblem(w, nethttp.StatusUnauthorized, map[string]any{
					"detail": "Missing required header: " + name,
				})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func HeaderToContext(
	headerName string,
	injector func(context.Context, string) (context.Context, error),
	failStatus int,
	failProb any, // Pure any payload supplied entirely by the downstream orchestrator
) Middleware {
	return func(next nethttp.Handler) nethttp.Handler {
		return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
			val := r.Header.Get(headerName)
			ctx, err := injector(r.Context(), val)
			if err != nil {
				WriteProblem(w, failStatus, failProb)
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
