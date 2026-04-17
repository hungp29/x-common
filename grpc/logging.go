package grpc

// Logger is a minimal logging surface for logadapter and tests.
type Logger interface {
	Debug(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	With(fields ...Field) Logger
}

type Field struct {
	Key   string
	Value interface{}
}
