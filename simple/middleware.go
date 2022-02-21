package simple

import (
	"fmt"
	"log"

	"github.com/Leadjet/logger"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (w *Logger) EchoMiddleware(_ *logger.WLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}

			// Retrieve email
			var email string
			if u := c.Get("user"); u != nil {
				t := u.(*jwt.Token)
				claims := t.Claims.(jwt.MapClaims)
				email = claims["email"].(string)
			}

			e := ""
			if email != "" {
				e = "(" + email + ")"
			}
			txt := fmt.Sprintf("%s %d %s %s", req.Method, res.Status, req.URL.Path, e)
			if err != nil {
				txt += "\n\t\t" + err.Error()
			}

			log.Println(txt)

			return nil
		}
	}
}
