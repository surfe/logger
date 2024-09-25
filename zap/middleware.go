package zap

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/surfe/logger/v2/key"
)

// EchoMiddleware is the echo middleware for Zap Logger
func (w *Logger) EchoMiddleware() echo.MiddlewareFunc {
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

			fields := []any{}
			appendFilledFieldsOnly(&fields, "remote_ip", c.RealIP())
			appendFilledFieldsOnly(&fields, key.External, false)
			appendFilledFieldsOnly(&fields, key.Email, email)
			appendFilledFieldsOnly(&fields, key.CompanyKey, companyKey)
			appendFilledFieldsOnly(&fields, key.ProcessingTime, time.Since(start).Milliseconds())
			appendFilledFieldsOnly(&fields, key.Method, req.Method)
			appendFilledFieldsOnly(&fields, key.URI, req.RequestURI)
			appendFilledFieldsOnly(&fields, key.Path, req.URL.Path)
			appendFilledFieldsOnly(&fields, key.Status, res.Status)
			appendFilledFieldsOnly(&fields, key.UserAgent, req.UserAgent())
			appendFilledFieldsOnly(&fields, key.APIVersion, req.Header.Get("API-Version"))
			appendFilledFieldsOnly(&fields, key.ExtensionVersion, req.Header.Get("Extension-Version"))

			if correlationID, isOk := req.Context().Value(key.CtxCorrelationID).(string); isOk && correlationID != "" {
				fields = append(fields, key.CorrelationID, correlationID)
			}

			log := w.Ctx(c.Request().Context()).With(fields...)
			n := res.Status
			// Status 5XX is logged as error as this should not happen in production.
			if n >= 500 {
				log.Err(err).Error("CRM Error")
			} else {
				log.Info("Incoming request")
			}

			return nil
		}
	}
}
