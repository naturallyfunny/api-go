package time

import (
	nethttp "net/http"

	apihttp "go.naturallyfunny.dev/api/http"
)

func HTTPWithZone(next nethttp.Handler) nethttp.Handler {
	return apihttp.HeaderToContext(
		"time-zone",
		NewContext,
		nethttp.StatusBadRequest,
		map[string]any{
			"detail": "Missing or invalid timezone header: expected a valid IANA timezone name (e.g. Asia/Jakarta)",
		},
	)(next)
}
