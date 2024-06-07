package zap

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/surfe/logger/key"
	"github.com/surfe/logger/logi"
)

// EchoMiddleware is the echo middleware for Zap Logger
func (w *Logger) EchoMiddleware(l logi.WLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			if l.MatchesAnyDiscardRule(res.Status, req.RequestURI) {
				return nil
			}

			var email string
			var companyKey string
			if u := c.Get("user"); u != nil {
				t := u.(*jwt.Token)
				claims := t.Claims.(jwt.MapClaims)
				email = claims["email"].(string)
				companyKey = claims["companyKey"].(string)
			}

			fields := []any{
				"remote_ip", c.RealIP(),
				key.External, false,
				key.Email, email,
				key.CompanyKey, companyKey,
				key.Latency, time.Since(start).Milliseconds(),
				key.Method, req.Method,
				key.URI, req.RequestURI,
				key.Path, req.URL.Path,
				key.Status, res.Status,
				key.UserAgent, req.UserAgent(),
				key.APIVersion, res.Header().Get("API-Version"),
				key.ExtensionVersion, res.Header().Get("Extension-Version"),
			}

			if correlationID, isOk := req.Context().Value(key.CtxCorrelationID).(string); isOk && correlationID != "" {
				fields = append(fields, key.CorrelationID, correlationID)
			}

			log := w.With(c.Request().Context(), fields...)
			n := res.Status
			switch {
			case n >= 500:
				log.Err(err).Error("CRM Error")
			case n >= 400:
				log.Err(err).Warn("Client error")
			case n >= 300:
				log.Info("Redirection")
			default:
				log.Info("Success")
			}

			return nil
		}
	}
}
