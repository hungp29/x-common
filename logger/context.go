package logger

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const loggerKey contextKey = "logger"

var (
	defaultFallback *zap.Logger
	fallbackOnce    sync.Once
)

func defaultOutLogger() *zap.Logger {
	fallbackOnce.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		l, err := cfg.Build()
		if err != nil {
			defaultFallback = zap.NewNop()
			return
		}
		defaultFallback = l
	})
	return defaultFallback
}

// With attaches a request-scoped logger to the context.
func With(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

// From returns the logger from ctx, or a JSON info-level stdout logger if none was set.
func From(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(loggerKey).(*zap.Logger); ok && l != nil {
		return l
	}
	return defaultOutLogger()
}
