package user

import (
	nethttp "net/http"

	"go.naturallyfunny.dev/api"
	apihttp "go.naturallyfunny.dev/api/http"
)

func HTTPMiddleware(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"user-id",
		NewContextWithID,
		api.NewError(api.Unauthenticated, "Missing required identity header"),
	)(next)
}
