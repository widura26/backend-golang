package routes

import (
	"backend-golang/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UserRoutes(r *echo.Group, jwtConfig echojwt.Config) {
	auth := r.Group("/auth")
	auth.POST("/signup", controllers.Signup)
	auth.POST("/signin", controllers.Signin)

	user := r.Group("/user")
	user.Use(echojwt.WithConfig(jwtConfig))
	user.GET("", controllers.GetUser)
}
