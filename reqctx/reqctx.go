package reqctx

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"github.com/surfe/logger/key"
)

func ContextMiddleware(apiVersion string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var email string
			var companyKey string
			if u := c.Get("user"); u != nil {
				t := u.(*jwt.Token)
				claims := t.Claims.(jwt.MapClaims)
				email = claims["email"].(string)
				companyKey = claims["companyKey"].(string)
			}

			// Extract data from headers if no user.
			if email == "" {
				email = c.Request().Header.Get(key.HeaderEmail)
			}
			if companyKey == "" {
				companyKey = c.Request().Header.Get(key.HeaderCompanyKey)
			}

			correlationID := c.Request().Header.Get(key.HeaderCorrelationID)
			if correlationID == "" {
				correlationID = random.String(32)
			}

			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key.CtxEmail, email)))
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key.CtxCompany, companyKey)))
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key.CtxCorrelationID, correlationID)))
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key.CtxAPIVersion, apiVersion)))

			c.Response().Header().Set(key.HeaderCorrelationID, correlationID)

			return next(c)
		}
	}
}
