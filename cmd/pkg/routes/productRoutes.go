package routes

import (
	"github.com/F1sssss/goecom/cmd/pkg/handlers"
	"github.com/F1sssss/goecom/cmd/pkg/services"

	"github.com/F1sssss/goecom/cmd/pkg/middleware"
	"github.com/labstack/echo"
)

// Init initializes all routes
func InitProductRoutes(e *echo.Echo) {

	protectedGroup := e.Group("/products", middleware.IsAuthorized)
	//protectedGroup.Use(middleware.HasRole("admin"))

	e.GET("/products", handlers.GetProducts)
	e.GET("/products/:id", handlers.GetProduct)
	protectedGroup.POST("", handlers.CreateProduct)
	protectedGroup.PATCH("/:id", handlers.UpdateProduct)
	protectedGroup.DELETE("/:id", handlers.DeleteProduct)
	protectedGroup.POST("/:id/shoppingcart", services.AddToCart)
	protectedGroup.GET("/shoppingcart", services.GetAllFromShoppingCart)
	protectedGroup.DELETE("/:id/shoppingcart", services.RemoveProductFromCart)

}
