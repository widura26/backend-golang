package routes

import (
	"backend-golang/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Group untuk auth
	auth := e.Group("/auth")
	auth.GET("/users", controllers.GetUser)
	auth.POST("/create-user", controllers.CreateUser)

	// Group untuk user
	// user := e.Group("/users")
	// user.GET("", controllers.GetUsers)
	// user.GET("/:id", controllers.GetUserByID)
	// user.POST("", controllers.CreateUser)
}
