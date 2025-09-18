package routes

import (
	"backend-golang/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func PostRoutes(r *echo.Group, jwtConfig echojwt.Config) {
	post := r.Group("/post")
	post.Use(echojwt.WithConfig(jwtConfig))
	post.POST("/create", controllers.CreatePost)
}
