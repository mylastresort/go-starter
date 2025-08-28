package internal

import (
	"log"
	"net/http"
	"server/internal/services"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Server *echo.Echo

func Init() {
	services.LoadConfig()
	services.LoadDatabase()
	LoadServer()
}

func LoadServer() {
	log.Default().Print("LOADING Server")
	Server = echo.New()
	Server.Use(middleware.Logger())
	AddUserRouter(Server)
}

func StartServer() {
	log.Default().Print("STARTING SERVER")
	if err := Server.Start(":" + strconv.Itoa(services.Conf.HTTP.PORT)); err != http.ErrServerClosed {
		Server.Logger.Fatal(err)
	}
}

func StopServer() {
}
