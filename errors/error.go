package errors

import "fmt"

type Code string

const (
	CodeInternal          Code = "INTERNAL"
	CodeNotFound          Code = "NOT_FOUND"
	CodeInvalidInput      Code = "INVALID_INPUT"
	CodeUnauthorized      Code = "UNAUTHORIZED"
	CodeForbidden         Code = "FORBIDDEN"
	CodeConflict          Code = "CONFLICT"
	CodeResourceExhausted Code = "RESOURCE_EXHAUSTED"
	CodePermissionDenied  Code = "PERMISSION_DENIED"
)

type AppError struct {
	Code    Code
	Reason  string // optional reason for the error
	Message string
	Err     error // original error (optional)
}

func (e *AppError) Error() string {
	return e.Message
}

// Unwrap returns the wrapped error, if any, for errors.Is / errors.As.
func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code Code, reason, msg string) *AppError {
	return &AppError{
		Code:    code,
		Reason:  reason,
		Message: msg,
	}
}

func Newf(code Code, reason, format string, args ...any) *AppError {
	return &AppError{
		Code:    code,
		Reason:  reason,
		Message: fmt.Sprintf(format, args...),
	}
}

func Wrap(code Code, reason, msg string, err error) *AppError {
	return &AppError{
		Code:    code,
		Reason:  reason,
		Message: msg,
		Err:     err,
	}
}

func Wrapf(code Code, reason, format string, err error, args ...any) *AppError {
	return &AppError{
		Code:    code,
		Reason:  reason,
		Message: fmt.Sprintf(format, args...),
		Err:     err,
	}
}
