package errors

import (
	"context"
	"errors"

	applogger "github.com/hungp29/x-common/logger"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
		st := status.New(mapCode(appErr.Code), appErr.Message)
		info := &errdetails.ErrorInfo{
			Reason: appErr.Reason,
		}

		st, _ = st.WithDetails(info)
		return st.Err()
	}

	// unknown error
	applogger.Error(context.Background(), "unknown error", "error", err.Error())
	// return generic internal error to avoid leaking details
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
	case CodeResourceExhausted:
		return codes.ResourceExhausted
	case CodePermissionDenied:
		return codes.PermissionDenied
	default:
		return codes.Internal
	}
}
