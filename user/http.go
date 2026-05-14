package user

import (
	nethttp "net/http"

	apihttp "go.naturallyfunny.dev/api/http"
)

func HTTPWithID(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"user-id",
		ContextWithID,
		nethttp.StatusUnauthorized,
		map[string]any{
			"detail": "Missing required identity header",
		},
	)(next)
}
