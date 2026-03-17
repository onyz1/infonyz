package infonyz

import "context"

type contextKey struct{}

var loggerKey = contextKey{}

// WithLogger returns a new context with the provided logger.
func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

// FromContext retrieves the logger from the context. If no logger is found, it returns a NoopLogger.
func FromContext(ctx context.Context) Logger {
	if log, ok := ctx.Value(loggerKey).(Logger); ok {
		return log
	}
	return NoopLogger()
}
