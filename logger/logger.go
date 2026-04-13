package logger

import (
	"context"

	"go.uber.org/zap"
)

func kvsToFields(args ...any) []zap.Field {
	if len(args) == 0 {
		return nil
	}
	if len(args)%2 != 0 {
		return append(kvsToFields(args[:len(args)-1]...), zap.Any("EXTRA_VALUE", args[len(args)-1]))
	}
	fs := make([]zap.Field, 0, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		k, _ := args[i].(string)
		fs = append(fs, zap.Any(k, args[i+1]))
	}
	return fs
}

func Info(ctx context.Context, msg string, args ...any) {
	From(ctx).Info(msg, kvsToFields(args)...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	From(ctx).Warn(msg, kvsToFields(args)...)
}

func Error(ctx context.Context, msg string, args ...any) {
	From(ctx).Error(msg, kvsToFields(args)...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	From(ctx).Debug(msg, kvsToFields(args)...)
}
