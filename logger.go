package logger

// Log is a package level variable, access logging functionality through "Log"
var Log Logger

// Basic fields
const (
	Email      string = "email"
	CompanyKey string = "company_key"
	Latency    string = "latency"
	Method     string = "method"
	URI        string = "uri"
	Status     string = "status"
	UserAgent  string = "user_agent"
	APIVersion string = "api_version"
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
}

// SetLogger is the setter for `Log` variable
func SetLogger(newLogger Logger) {
	Log = newLogger
}
