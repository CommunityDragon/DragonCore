package auth_oauth2

import (
	"dragonback/lib/models/controller"
	"dragonback/lib/models/swaggerdocs"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var handler = controller.New("/oauth2").

	// registers a user
	Handler(
		controller.POST,
		"/register",
		swaggerdocs.New().
			SetDescription("Registers a new user.").
			AddParamBody(loginPayload{}, "body", "Login Payload", true).
			AddResponse(http.StatusOK, "successful", []string{}, nil),
		func(c echo.Context) (err error) {
			payload := new(loginPayload)
			if err = c.Bind(payload); err != nil {
				return
			}

			if err = c.Validate(payload); err != nil {
				fmt.Print(err)
				return
			}
			return c.JSON(http.StatusOK, payload)
	}).

	// logs a user in
	Handler(controller.POST, "/login", nil, func(c echo.Context) (err error) {
		payload := new(loginPayload)
		if err = c.Bind(payload); err != nil {
			return
		}
		return c.JSON(http.StatusOK, nil)
	})