package timezone

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type key struct{}

var ContextKey = key{}

func NewContext(ctx context.Context, tz string) (context.Context, error) {
	if tz == "" {
		return ctx, errors.New("timezone cannot be empty")
	}
	if _, err := time.LoadLocation(tz); err != nil {
		return ctx, fmt.Errorf("invalid IANA timezone %q: %w", tz, err)
	}
	return context.WithValue(ctx, ContextKey, tz), nil
}

func FromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(ContextKey).(string)
	if !ok || val == "" {
		return "", errors.New("timezone not found in context")
	}
	return val, nil
}
