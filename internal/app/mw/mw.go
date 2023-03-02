package mw

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		headerValue := context.Request().Header.Get("User-Role")

		if strings.EqualFold(headerValue, roleAdmin) {
			log.Println("red button user detected")
		}

		if err := next(context); err != nil {
			return err
		}

		return nil
	}
}
