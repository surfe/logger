package zap

import (
	"time"

	"github.com/Leadjet/logger"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// EchoMiddleware is the echo middleware for Zap Logger
func (w *Logger) EchoMiddleware(l *logger.WLogger) echo.MiddlewareFunc {
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
				logger.EmailKey, email,
				logger.CompanyKey, companyKey,
				logger.LatencyKey, time.Since(start).String(),
				logger.MethodKey, req.Method,
				logger.URIKey, req.RequestURI,
				logger.StatusKey, res.Status,
				logger.UserAgentKey, req.UserAgent(),
				logger.APIVersionKey, req.Header.Get("X-API-Version"),
			}

			n := res.Status
			switch {
			case n >= 500:
				w.Errorw("CRM Error", err, fields...)
			case n >= 400:
				w.Errorw("Server Error", err, fields...)
			case n >= 300:
				w.Infow("Redirection", fields...)
			default:
				w.Infow("Success", fields...)
			}

			return nil
		}
	}
}
