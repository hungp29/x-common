package logger

import (
	"context"
	"sync"

	appcontext "github.com/hungp29/x-common/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultFallback *zap.Logger
	fallbackOnce    sync.Once
)

func defaultOutLogger() *zap.Logger {
	fallbackOnce.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		cfg.DisableCaller = true
		l, err := cfg.Build()
		if err != nil {
			defaultFallback = zap.NewNop()
			return
		}
		defaultFallback = l
	})
	return defaultFallback
}

// From returns the logger from ctx, or a JSON info-level stdout logger if none was set.
func From(ctx context.Context) *zap.Logger {
	if l, ok := appcontext.Logger(ctx); ok && l != nil {
		return l
	}
	return defaultOutLogger()
}
