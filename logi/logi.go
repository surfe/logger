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
	// WithContext will return logger with filled basic fields (email, company, corelationID)
	WithContext(ctx context.Context) Logger

	// Errorf logs a templated message with the provided error
	Errorf(format string, err interface{}, args ...interface{})

	// Errorw logs a message with optional fields
	Errorw(msg string, err interface{}, keysAndValues ...interface{})

	// Error logs a simple message with the provided error
	Error(err interface{}, args ...interface{})

	// Infof logs a templated message with optional fields
	Infof(format string, args ...interface{})

	// Infow logs a message with optional fields
	Infow(msg string, keysAndValues ...interface{})

	// Info logs a simple message
	Info(args ...interface{})

	// Debugf logs a templated message with optional fields
	Debugf(format string, args ...interface{})

	// Debugw logs a message with optional fields
	Debugw(msg string, keysAndValues ...interface{})

	// Debug logs a simple message
	Debug(args ...interface{})

	// Sync cleanups before exiting
	Sync()

	// EchoMiddleware returns EchoMiddleware of current logger
	EchoMiddleware(WLogger) echo.MiddlewareFunc
}
