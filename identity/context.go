package identity

import (
	"context"
	"errors"
)

type contextKey string

const (
	UserIDKey    contextKey = "user_id"
	SessionIDKey contextKey = "session_id"
)

func NewContextWithUserID(ctx context.Context, userID string) (context.Context, error) {
	if userID == "" {
		return ctx, errors.New("user ID cannot be empty")
	}
	return context.WithValue(ctx, UserIDKey, userID), nil
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(UserIDKey).(string)
	if !ok || val == "" {
		return "", errors.New("user ID not found in context")
	}
	return val, nil
}

func NewContextWithSessionID(ctx context.Context, sessionID string) (context.Context, error) {
	if sessionID == "" {
		return ctx, errors.New("session ID cannot be empty")
	}
	return context.WithValue(ctx, SessionIDKey, sessionID), nil
}

func GetSessionIDFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(SessionIDKey).(string)
	if !ok || val == "" {
		return "", errors.New("session ID not found in context")
	}
	return val, nil
}
