package mw

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"reflect"
)

// PayloadBinder binds the payload to the context
func PayloadBinder(payload interface{}) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			pl := reflect.New(reflect.ValueOf(payload).Type()).Interface()
			if err = c.Bind(pl); err != nil {
				return
			}
			if err = c.Validate(pl); err != nil {
				fmt.Print(err)
				return
			}
			c.Set("payload", pl)
			return next(c)
		}
	}
}

