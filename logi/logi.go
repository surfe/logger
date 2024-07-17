package logi

import (
	"context"

	"github.com/labstack/echo/v4"
)

// Logger represents common interface for logging functionality
type Logger interface {
	// With returns logger with basic fields based on context and your custom fields
	With(ctx context.Context, keysAndValues ...any) Logger

	// Err returns logger with the provided error
	Err(err error) Logger

	// Errorf logs a templated message
	Errorf(format string, args ...any)

	// Error logs a simple message
	Error(args ...any)

	// Warnf logs a templated message
	Warnf(format string, args ...any)

	// Warn logs a simple message
	Warn(args ...any)

	// Infof logs a templated message
	Infof(format string, args ...any)

	// Info logs a simple message
	Info(args ...any)

	// Debugf logs a templated message
	Debugf(format string, args ...any)

	// Debug logs a simple message
	Debug(args ...any)

	// Fatalf logs a templated message, then exits
	Fatalf(format string, args ...any)

	// Fatal logs a simple message, then exits
	Fatal(args ...any)

	// Sync cleanups before exiting
	Sync()

	// EchoMiddleware returns EchoMiddleware of current logger
	EchoMiddleware() echo.MiddlewareFunc
}
