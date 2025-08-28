package internal

import (
	"server/internal/controllers"

	"github.com/labstack/echo"
)

func AddUserRouter(e *echo.Echo) {
	rootRoute := e.Group("/")
	rootRoute.GET("/", controllers.GetUsers)
}
