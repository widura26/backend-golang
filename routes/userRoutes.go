package routes

import (
	"backend-golang/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.POST("/signup", controllers.Signup)
	auth.POST("/signin", controllers.Signin)
}

func PostRoutes(r *echo.Group) {
	r.GET("/user", controllers.GetUser)
}
