package session

import (
	"context"
	"errors"
)

type key struct{}

var contextKey key

func NewContextWithID(ctx context.Context, id string) (context.Context, error) {
	if id == "" {
		return ctx, errors.New("session ID cannot be empty")
	}
	return context.WithValue(ctx, contextKey, id), nil
}

func IDFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(contextKey).(string)
	if !ok || val == "" {
		return "", errors.New("session ID not found in context")
	}
	return val, nil
}
