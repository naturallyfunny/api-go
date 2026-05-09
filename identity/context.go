package identity

import (
	"context"
	"errors"
)

type contextKey string

const (
	userIDKey    contextKey = "user_id"
	sessionIDKey contextKey = "session_id"
)

func NewContextWithUserID(ctx context.Context, userID string) (context.Context, error) {
	if userID == "" {
		return ctx, errors.New("user ID cannot be empty")
	}
	return context.WithValue(ctx, userIDKey, userID), nil
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(userIDKey).(string)
	if !ok || val == "" {
		return "", errors.New("user ID not found in context")
	}
	return val, nil
}

func NewContextWithSessionID(ctx context.Context, sessionID string) (context.Context, error) {
	if sessionID == "" {
		return ctx, errors.New("session ID cannot be empty")
	}
	return context.WithValue(ctx, sessionIDKey, sessionID), nil
}

func GetSessionIDFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(sessionIDKey).(string)
	if !ok || val == "" {
		return "", errors.New("session ID not found in context")
	}
	return val, nil
}
