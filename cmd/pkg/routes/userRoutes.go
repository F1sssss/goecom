package routes

import (
	"github.com/F1sssss/goecom/cmd/pkg/handlers"
	"github.com/labstack/echo"
)

// Init initializes all routes
func InitUserRoutes(e *echo.Echo) {

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.PATCH("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

}
