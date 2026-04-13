package errors

import "fmt"

type Code string

const (
	CodeInternal     Code = "INTERNAL"
	CodeNotFound     Code = "NOT_FOUND"
	CodeInvalidInput Code = "INVALID_INPUT"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeForbidden    Code = "FORBIDDEN"
	CodeConflict     Code = "CONFLICT"
	CodeResourceExhausted Code = "RESOURCE_EXHAUSTED"
)

type AppError struct {
	Code    Code
	Message string
	Err     error // original error (optional)
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code Code, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

func Newf(code Code, format string, args ...any) *AppError {
	return &AppError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

func Wrap(code Code, msg string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func Wrapf(code Code, format string, err error, args ...any) *AppError {
	return &AppError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
		Err:     err,
	}
}
