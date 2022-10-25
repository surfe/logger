package zap

import (
	"time"

	"github.com/Leadjet/logger/key"
	"github.com/Leadjet/logger/logi"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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

			fields := []interface{}{
				"remote_ip", c.RealIP(),
				key.Email, email,
				key.CompanyKey, companyKey,
				key.Latency, time.Since(start).String(),
				key.Method, req.Method,
				key.URI, req.RequestURI,
				key.Status, res.Status,
				key.UserAgent, req.UserAgent(),
				key.APIVersion, req.Header.Get("X-API-Version"),
			}

			if corelationID, isOk := req.Context().Value(key.CtxCorelationID).(string); isOk {
				fields = append(fields, key.CorelationID, corelationID)
			}

			n := res.Status
			switch {
			case n >= 500:
				w.Errorw("CRM Error", err, fields...)
			case n >= 400:
				w.Warnw("Server Error", err, fields...)
			case n >= 300:
				w.Infow("Redirection", fields...)
			default:
				w.Infow("Success", fields...)
			}

			return nil
		}
	}
}
