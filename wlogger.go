package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/surfe/logger/v2/logi"
)

type wLogger struct{}

// Use initiates a wLogger by wrapping provided Logger instance, and sets global logger variable
func Use(l logi.Logger) *wLogger {
	logger = l

	return &wLogger{}
}

// EchoMiddleware is the getter for wrapped logger's EchoMiddleware
func (l *wLogger) EchoMiddleware() echo.MiddlewareFunc {
	return logger.EchoMiddleware()
}
