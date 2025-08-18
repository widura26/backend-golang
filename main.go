package main

import (
	"backend-golang/config"
	"backend-golang/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", controllers.GetUser)
	e.POST("/create-user", controllers.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
