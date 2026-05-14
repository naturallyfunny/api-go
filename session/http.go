package session

import (
	nethttp "net/http"

	apihttp "go.naturallyfunny.dev/api/http"
)

func HTTPWithID(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"session-id",
		ContextWithID,
		nethttp.StatusUnauthorized,
		map[string]any{
			"detail": "Missing required session identity header",
		},
	)(next)
}
