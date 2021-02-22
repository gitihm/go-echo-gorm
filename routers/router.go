package routers

import (
	"main/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	api := e.Group("/api")
	users := api.Group("/users")

	// Users Controller
	users.GET("", controllers.GetUsers)

	return e
}
