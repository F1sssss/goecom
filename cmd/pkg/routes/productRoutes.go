package routes

import (
	"github.com/F1sssss/goecom/cmd/pkg/handlers"
	"github.com/F1sssss/goecom/cmd/pkg/middleware"
	"github.com/labstack/echo"
)

// Init initializes all routes
func InitProductRoutes(e *echo.Echo) {

	protectedGroup := e.Group("/products", middleware.IsAuthorized)

	e.GET("/products", handlers.GetProducts)
	protectedGroup.GET("/:id", handlers.GetProduct)
	protectedGroup.POST("", handlers.CreateProduct)
	protectedGroup.PATCH("/:id", handlers.UpdateProduct)
	protectedGroup.DELETE("/:id", handlers.DeleteProduct)

}
