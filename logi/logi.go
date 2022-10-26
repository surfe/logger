package logi

import (
	"context"

	"github.com/labstack/echo/v4"
)

type WLogger interface {
	MatchesAnyDiscardRule(status int, uri string) bool
}

// Logger represents common interface for logging functionality
type Logger interface {
	// With returns logger with basic fields based on context and your custom fields
	With(ctx context.Context, keysAndValues ...any) Logger

	// Err is a shorter version of With(ctx, "error", err)
	Err(err error) Logger

	// Errorf logs a templated message with the provided error
	Errorf(format string, args ...any)

	// Error logs a simple message with the provided error
	Error(args ...any)

	// Warnf logs a templated message with the provided error
	Warnf(format string, args ...any)

	// Warn logs a simple message with the provided error
	Warn(args ...any)

	// Infof logs a templated message with optional fields
	Infof(format string, args ...any)

	// Info logs a simple message
	Info(args ...any)

	// Debugf logs a templated message with optional fields
	Debugf(format string, args ...any)

	// Debug logs a simple message
	Debug(args ...any)

	// Sync cleanups before exiting
	Sync()

	// EchoMiddleware returns EchoMiddleware of current logger
	EchoMiddleware(WLogger) echo.MiddlewareFunc
}
