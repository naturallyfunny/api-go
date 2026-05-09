package identity

import (
	"net/http"

	"go.naturallyfunny.dev/api"
)

func WithUserID(next http.Handler) http.Handler {
	return api.HeaderToContext(
		"x-user-id",
		NewContextWithUserID,
		api.NewError(api.Unauthenticated, "Missing required identity header"),
	)(next)
}

func WithSessionID(next http.Handler) http.Handler {
	return api.HeaderToContext(
		"x-session-id",
		NewContextWithSessionID,
		api.NewError(api.Unauthenticated, "Missing required session identity header"),
	)(next)
}

func WithTimezone(next http.Handler) http.Handler {
	return api.HeaderToContext(
		"x-user-timezone",
		NewContextWithTimezone,
		api.NewError(api.InvalidArgument, "Missing or invalid timezone header: expected a valid IANA timezone name (e.g. Asia/Jakarta)"),
	)(next)
}
