package grpc

import (
	"context"

	gogrpc "google.golang.org/grpc"
)

func UnaryContextInjectorInterceptor(
	injector func(ctx context.Context) context.Context,
) gogrpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *gogrpc.UnaryServerInfo,
		handler gogrpc.UnaryHandler,
	) (interface{}, error) {

		ctx = injector(ctx)
		return handler(ctx, req)
	}
}
