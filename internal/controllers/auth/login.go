package auth

import (
	"net/http"
	"server/internal/services"

	"github.com/labstack/echo/v4"
)

type LoginUserType struct {
	Email string `validate:"required"`
	// Password string `validate:"required"`
}

func Login(c echo.Context) error {
	var newUser LoginUserType

	err := c.Bind(&newUser)
	if err != nil {
		return echo.ErrBadRequest
	}

	err = services.ValidateStruct(newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := services.GetUserByEmail(newUser.Email)
	if err != nil {
		return echo.ErrNotFound
	}

	response, err := RevokeToken(user, "")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
