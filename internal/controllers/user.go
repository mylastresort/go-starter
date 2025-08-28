package controllers

import (
	"net/http"
	"server/internal/services"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	users, err := services.GetUsers()
	if err != nil {
		response := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusOK, response)
	}

	response := map[string]interface{}{
		"data": users,
	}

	return c.JSON(http.StatusOK, response)
}
