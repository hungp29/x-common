package grpc

import (
	"context"

	"go.uber.org/zap"
)

type contextKey string

const (
	ContextKeyUserID        contextKey = "user_id"
	ContextKeyCorrelationID contextKey = "correlation_id"
	ContextKeyLogger        contextKey = "logger"
)

// WithUserID attaches the authenticated user id for downstream handlers.
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextKeyUserID, userID)
}

// UserID returns the user id previously stored with WithUserID.
func UserID(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(ContextKeyUserID).(string)
	return v, ok && v != ""
}

// WithCorrelationID attaches a correlation / trace id for logging across calls.
func WithCorrelationID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ContextKeyCorrelationID, id)
}

// CorrelationID returns the correlation id from context, if set.
func CorrelationID(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(ContextKeyCorrelationID).(string)
	return v, ok && v != ""
}

// WithLogger attaches a logger to the context.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ContextKeyLogger, logger)
}

// Logger returns the logger from the context, if set.
func Logger(ctx context.Context) (*zap.Logger, bool) {
	v, ok := ctx.Value(ContextKeyLogger).(*zap.Logger)
	return v, ok && v != nil
}
