package time

import (
	nethttp "net/http"

	"go.naturallyfunny.dev/api"
	apihttp "go.naturallyfunny.dev/api/http"
)

func HTTPWithZone(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"time-zone",
		NewContext,
		api.NewError(api.InvalidArgument, "Missing or invalid timezone header: expected a valid IANA timezone name (e.g. Asia/Jakarta)"),
	)(next)
}
