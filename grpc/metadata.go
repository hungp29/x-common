package grpc

import (
	"context"

	appcontext "github.com/hungp29/x-common/context"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryMetadataInterceptor(opts Options) gogrpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *gogrpc.UnaryServerInfo,
		handler gogrpc.UnaryHandler,
	) (interface{}, error) {

		md, _ := metadata.FromIncomingContext(ctx)

		var correlationID string
		var userID string

		if vals := md.Get(opts.CorrelationIDHeader); len(vals) > 0 {
			correlationID = vals[0]
		} else {
			correlationID = opts.GenerateCorrelationID()
		}

		if vals := md.Get(opts.UserIDHeader); len(vals) > 0 {
			userID = vals[0]
		}

		ctx = appcontext.WithCorrelationID(ctx, correlationID)

		if opts.EnableUserIDContext && userID != "" {
			ctx = appcontext.WithUserID(ctx, userID)
		}

		return handler(ctx, req)
	}
}
