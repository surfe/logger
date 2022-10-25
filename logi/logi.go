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
	Errorf(format string, err error, args ...interface{})

	// Errorw logs a message with optional fields
	Errorw(msg string, err error, keysAndValues ...interface{})

	// Error logs a simple message with the provided error
	Error(err error, args ...interface{})

	// Warnf logs a templated message with the provided error
	Warnf(format string, err error, args ...interface{})

	// Warnw logs a message with optional fields
	Warnw(msg string, err error, keysAndValues ...interface{})

	// Warn logs a simple message with the provided error
	Warn(err error, args ...interface{})

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
