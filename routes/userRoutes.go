package routes

import (
	"backend-golang/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.GET("/user", controllers.GetUser)
	auth.POST("/signup", controllers.Signup)
	auth.POST("/signin", controllers.Signin)
}
