package identity

import (
	nethttp "net/http"

	"go.naturallyfunny.dev/api"
	apihttp "go.naturallyfunny.dev/api/http"
)

func HTTPWithUserID(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"x-user-id",
		NewContextWithUserID,
		api.NewError(api.Unauthenticated, "Missing required identity header"),
	)(next)
}

func HTTPWithSessionID(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"x-session-id",
		NewContextWithSessionID,
		api.NewError(api.Unauthenticated, "Missing required session identity header"),
	)(next)
}

func HTTPWithUserTimezone(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"x-user-timezone",
		NewContextWithTimezone,
		api.NewError(api.InvalidArgument, "Missing or invalid timezone header: expected a valid IANA timezone name (e.g. Asia/Jakarta)"),
	)(next)
}
