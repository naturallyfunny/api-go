package session

import (
	nethttp "net/http"

	"go.naturallyfunny.dev/api"
	apihttp "go.naturallyfunny.dev/api/http"
)

func HTTPWithID(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"session-id",
		NewContextWithID,
		api.NewError(api.Unauthenticated, "Missing required session identity header"),
	)(next)
}
