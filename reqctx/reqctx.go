package reqctx

import (
	"context"

	"github.com/Leadjet/logger/key"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
)

func ContextMiddleware() echo.MiddlewareFunc {
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

			corelationID := random.String(32)
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key.CtxEmail, email)))
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key.CtxCompany, companyKey)))
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key.CtxCorelationID, corelationID)))

			c.Response().Header().Set("x-corelation-id", corelationID)

			return next(c)
		}
	}
}
