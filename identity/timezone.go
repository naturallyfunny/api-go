package identity

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// timezoneKey is an unexported zero-size type. Using a distinct named type
// (rather than a plain string) as the context key prevents collision even if
// another package coincidentally uses the same string value.
type timezoneKey struct{}

// TimezoneKey is exported so integrators can bridge timezone propagation to
// other packages (e.g. via a WithTimezoneFromContext(key any) option) without
// those packages needing to import this one. Because timezoneKey is unexported,
// external packages cannot construct an instance of it — making this var
// effectively immutable from outside this package.
var TimezoneKey = timezoneKey{}

func NewContextWithTimezone(ctx context.Context, tz string) (context.Context, error) {
	if tz == "" {
		return ctx, errors.New("timezone cannot be empty")
	}
	if _, err := time.LoadLocation(tz); err != nil {
		return ctx, fmt.Errorf("invalid IANA timezone %q: %w", tz, err)
	}
	return context.WithValue(ctx, TimezoneKey, tz), nil
}

func GetTimezoneFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(TimezoneKey).(string)
	if !ok || val == "" {
		return "", errors.New("timezone not found in context")
	}
	return val, nil
}
