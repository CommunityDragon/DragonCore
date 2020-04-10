package auth_oauth2

import (
	mw "dragoncore/lib/middleware"
	"dragoncore/lib/models/controller"
	"dragoncore/lib/models/swaggerdocs"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var handler = controller.New("/auth").

	// registers a user
	Handler(
		controller.POST,
		"/register",

		swaggerdocs.New().
			SetDescription("Registers a new user.").
			AddParamBody(loginPayload{}, "body", "Login Payload", true).
			AddResponse(http.StatusCreated, "successful", userResponse{}, nil),

		registerUser,
		mw.PayloadBinder(loginPayload{}),
		middleware.Recover()).

	// registers a user
	Handler(
		controller.GET,
		"/register",

		swaggerdocs.New().
			SetDescription("Loads the registration page."),

		renderRegister,
		middleware.CSRF()).

	// logs a user in
	Handler(
		controller.POST,
		"/login",

		swaggerdocs.New().
			SetDescription("Logs in a user.").
			AddParamBody(loginPayload{}, "body", "Login Payload", true).
			AddResponse(http.StatusOK, "successful", userResponse{}, nil),

		loginUser,
		mw.PayloadBinder(loginPayload{}),
		middleware.Recover())