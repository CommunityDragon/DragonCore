package auth_oauth2

import (
	"dragonback/lib/db"
	"dragonback/lib/entities"
	"dragonback/lib/util"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"net/http"
)

// registerUser registers a user
func registerUser(c echo.Context) error {
	payload := c.Get("payload").(*loginPayload)
	password, err := util.HashPassword(payload.Password)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to process password. Try another one please")
	}

	if exists := entities.Users(qm.Where("email=?", payload.Email)).ExistsP(c.Request().Context(), db.DB()); exists {
		return echo.NewHTTPError(http.StatusBadRequest, "A user with this email already exists")
	}

	user := &entities.User{
		Email: payload.Email,
		Password: password,
	}

	user.InsertP(c.Request().Context(), db.DB(), boil.Infer())
	return c.JSON(http.StatusOK, &userResponse{user.ID, user.Email})
}

func loginUser(c echo.Context) (err error) {
	payload := c.Get("payload").(*loginPayload)
	user, err := entities.Users(qm.Where("email=?", payload.Email)).One(c.Request().Context(), db.DB())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	if match := util.VerifyPassword(payload.Password, user.Password); !match {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	return c.JSON(http.StatusOK, &userResponse{user.ID, user.Email})
}

