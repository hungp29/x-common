package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Map AppError -> gRPC status
func ToGRPC(err error) error {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return status.Error(mapCode(appErr.Code), appErr.Message)
	}

	// unknown error
	return status.Error(codes.Internal, "internal server error")
}

func mapCode(code Code) codes.Code {
	switch code {
	case CodeNotFound:
		return codes.NotFound
	case CodeInvalidInput:
		return codes.InvalidArgument
	case CodeUnauthorized:
		return codes.Unauthenticated
	case CodeForbidden:
		return codes.PermissionDenied
	case CodeConflict:
		return codes.AlreadyExists
	default:
		return codes.Internal
	}
}