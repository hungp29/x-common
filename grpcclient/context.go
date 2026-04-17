package grpcclient

import "context"

type contextKey string

const (
	userIDKey        contextKey = "user_id"
	correlationIDKey contextKey = "correlation_id"
)

// WithUserID attaches a user id for the unary client interceptor to send as gRPC metadata.
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// UserIDFromContext returns the user id set by WithUserID, or "" if unset.
func UserIDFromContext(ctx context.Context) string {
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		return userID
	}
	return ""
}

// WithCorrelationID attaches a correlation id for the unary client interceptor.
func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationIDKey, correlationID)
}

// CorrelationIDFromContext returns the correlation id set by WithCorrelationID, or "" if unset.
func CorrelationIDFromContext(ctx context.Context) string {
	if correlationID, ok := ctx.Value(correlationIDKey).(string); ok {
		return correlationID
	}
	return ""
}
