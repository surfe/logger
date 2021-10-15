package zap

import (
	logger "github.com/Leadjet/Logger"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

// EchoMiddleware is the echo middleware for Zap logger
func EchoMiddleware(l logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

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
				logger.Email, email,
				logger.CompanyKey, companyKey,
				logger.Latency, time.Since(start).String(),
				logger.Method, req.Method,
				logger.URI, req.RequestURI,
				logger.Status, res.Status,
				logger.UserAgent, req.UserAgent(),
				logger.APIVersion, req.Header.Get("X-API-Version"),
			}

			n := res.Status
			switch {
			case n >= 500:
				l.Errorw("CRM Error", err, fields...)
			case n >= 400:
				l.Errorw("Server Error", err, fields...)
			case n >= 300:
				l.Infow("Redirection", fields...)
			default:
				l.Infow("Success", fields...)
			}

			return nil
		}
	}
}
