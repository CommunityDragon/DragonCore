package mw

import (
	"dragoncore/lib/services"
	"fmt"
	"github.com/labstack/echo/v4"
)

// InjectService injects a service
func InjectService(service services.Service) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			c.Set(fmt.Sprintf("service_%s", service.Name()), service)
			return next(c)
		}
	}
}
