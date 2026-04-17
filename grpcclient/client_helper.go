package grpcclient

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewClient dials a gRPC server at the given address and returns a grpc.ClientConn.
// The returned *grpc.ClientConn must be closed by the caller when the application shuts down.
func NewClient(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, fmt.Errorf("dial gRPC server at %q: %w", addr, err)
	}
	return conn, nil
}
