package grpc

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// ContextAttach binds the request-scoped *zap.Logger to ctx for downstream handlers.
// Wire with github.com/hungp29/x-common/logger.AttachZap when using the shared logger package.
// If nil, the interceptor does not attach a logger to context (handlers cannot use logger.From).
type ContextAttach func(ctx context.Context, log *zap.Logger) context.Context

type Options struct {
	UserIDHeader        string
	CorrelationIDHeader string

	GenerateCorrelationID func() string

	EnableUserIDContext bool

	ContextAttach ContextAttach
}

func DefaultOptions() Options {
	return Options{
		UserIDHeader:        "user_id",
		CorrelationIDHeader: "correlation_id",
		GenerateCorrelationID: func() string {
			return uuid.NewString()
		},
		EnableUserIDContext: true,
	}
}

// NormalizeOptions fills empty header / id-generator fields from DefaultOptions.
// If o is the zero value, it returns DefaultOptions() so defaults include EnableUserIDContext=true.
// Otherwise EnableUserIDContext is preserved as set on o.
func NormalizeOptions(o Options) Options {
	if o.CorrelationIDHeader == "" && o.UserIDHeader == "" && o.GenerateCorrelationID == nil && !o.EnableUserIDContext {
		return DefaultOptions()
	}
	d := DefaultOptions()
	if o.CorrelationIDHeader == "" {
		o.CorrelationIDHeader = d.CorrelationIDHeader
	}
	if o.UserIDHeader == "" {
		o.UserIDHeader = d.UserIDHeader
	}
	if o.GenerateCorrelationID == nil {
		o.GenerateCorrelationID = d.GenerateCorrelationID
	}
	return o
}
