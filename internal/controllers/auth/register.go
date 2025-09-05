package auth

import (
	"net/http"
	"server/internal/services"
	"server/internal/services/users"

	"github.com/labstack/echo/v4"
)

type RegisterUserType struct {
	Name     string `validate:"required,min=4"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,containsany=!@#?*"`
}

func Register(c echo.Context) error {
	var newUser RegisterUserType

	err := c.Bind(&newUser)
	if err != nil {
		return echo.ErrBadRequest
	}

	var user users.CreateUserType
	user.Provider = ""
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password

	err = services.ValidateStruct(newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = users.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "success")
}
