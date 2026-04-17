package grpc

import (
	"context"
	"strings"
	"time"

	appcontext "github.com/hungp29/x-common/context"
	"go.uber.org/zap"
	gogrpc "google.golang.org/grpc"
)

func UnaryLoggingInterceptor(baseLogger *zap.Logger) gogrpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *gogrpc.UnaryServerInfo,
		handler gogrpc.UnaryHandler,
	) (interface{}, error) {

		// parse method
		parts := strings.Split(info.FullMethod, "/")
		service, method := "unknown", "unknown"
		if len(parts) >= 3 {
			service, method = parts[1], parts[2]
		}

		// get correlation ID from context (set by UnaryMetadataInterceptor)
		correlationID, _ := appcontext.CorrelationID(ctx)

		logger := baseLogger.With(
			zap.String("grpc_service", service),
			zap.String("grpc_method", method),
			zap.String("correlation_id", correlationID),
		)

		// inject logger into context
		ctx = appcontext.WithLogger(ctx, logger)

		start := time.Now()

		logger.Debug("rpc call",
			zap.String("state", "start"),
		)

		resp, err := handler(ctx, req)

		duration := time.Since(start).Milliseconds()

		if err != nil {
			logger.Error("rpc call error",
				zap.Error(err),
				zap.Int64("duration_ms", duration),
			)
		} else {
			logger.Debug("rpc call",
				zap.String("state", "end"),
				zap.Bool("ok", true),
				zap.Int64("duration_ms", duration),
			)
		}

		return resp, err
	}
}
