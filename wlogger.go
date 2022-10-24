package logger

import (
	"github.com/Leadjet/logger/logi"
	"github.com/labstack/echo/v4"
)

type wLogger struct {
	DiscardRules []DiscardRule
}

// Use initiates a wLogger by wrapping provided Logger instance, and sets global logger variable
func Use(l logi.Logger) *wLogger {
	logger = l

	return &wLogger{}
}

// EchoMiddleware is the getter for wrapped logger's EchoMiddleware
func (l *wLogger) EchoMiddleware() echo.MiddlewareFunc {
	return logger.EchoMiddleware(l)
}

// MatchesAnyDiscardRule checks if provided status and uri params are matching with any discard rule
func (l *wLogger) MatchesAnyDiscardRule(status int, uri string) bool {
	for _, rule := range l.DiscardRules {
		if rule.Status == status {
			for _, rURI := range rule.URIs {
				if rURI == uri {
					return true
				}
			}
		}
	}

	return false
}

// DiscardRule defines struct which may used to discard specific logs
type DiscardRule struct {
	Status int
	URIs   []string
}
