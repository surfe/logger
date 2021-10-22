package logger

import (
	"github.com/Leadjet/logger/simple"
)

var log Logger

// Basic fields
const (
	EmailKey      string = "email"
	CompanyKey    string = "company_key"
	LatencyKey    string = "latency"
	MethodKey     string = "method"
	URIKey        string = "uri"
	StatusKey     string = "status"
	UserAgentKey  string = "user_agent"
	APIVersionKey string = "api_version"
	PayloadKey    string = "payload"
	UserKey       string = "user"
)

// Logger represent common interface for logging functionality
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
}

// Log is the getter for `log` variable which defaults to a wrapped simple `log` instance if `nil`
func Log() Logger {
	if log == nil {
		log = &simple.Logger{}
	}

	return log
}

// SetLogger is the setter for `Log` variable
func SetLogger(newLogger Logger) {
	log = newLogger
}
