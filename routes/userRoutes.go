package routes

import (
	"backend-golang/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Group untuk auth
	auth := e.Group("/auth")
	auth.GET("/user", controllers.GetUser)
	auth.POST("/signup", controllers.Signup)
	auth.POST("/signin", controllers.Signin)
	auth.GET("/update-user", controllers.UpdateUser)
}
