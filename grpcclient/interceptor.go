package grpcclient

import (
	"context"

	appcontext "github.com/hungp29/x-common/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryClientInterceptor propagates user_id and correlation_id from context into outgoing metadata.
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		userID, _ := appcontext.UserID(ctx)
		correlationID, _ := appcontext.CorrelationID(ctx)

		md := metadata.New(map[string]string{})
		if userID != "" {
			md.Set("user_id", userID)
		}
		if correlationID != "" {
			md.Set("correlation_id", correlationID)
		}

		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
