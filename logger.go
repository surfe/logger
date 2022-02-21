package logger

import (
	"github.com/labstack/echo/v4"
)

var logger Logger

// Log is the getter for global `logger` variable
func Log() Logger {
	return logger
}

// Logger represents common interface for logging functionality
type Logger interface {
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
	EchoMiddleware(*WLogger) echo.MiddlewareFunc
}

// Basic fields
const (
	EmailKey      = "email"
	CompanyKey    = "company_key"
	LatencyKey    = "latency"
	MethodKey     = "method"
	URIKey        = "uri"
	StatusKey     = "status"
	UserAgentKey  = "user_agent"
	APIVersionKey = "api_version"
	PayloadKey    = "payload"
	UserKey       = "user"
)
