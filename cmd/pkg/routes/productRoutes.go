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
	e.GET("/search/products", services.SearchProducts)
	protectedGroup.POST("", handlers.CreateProduct)
	protectedGroup.PATCH("/:id", handlers.UpdateProduct)
	protectedGroup.DELETE("/:id", handlers.DeleteProduct)
	protectedGroup.POST("/:id/shoppingcart", services.AddToCart)
	protectedGroup.GET("/shoppingcart", services.GetAllFromShoppingCart)
	protectedGroup.DELETE("/:id/shoppingcart", services.RemoveProductFromCart)

	// Product Reviews
	protectedGroup.GET("/:id/reviews", handlers.GetProductReviews)
	protectedGroup.POST("/:id/reviews", handlers.AddReview)
	protectedGroup.PATCH("/:id/reviews/", handlers.UpdateReview)
	protectedGroup.DELETE("/:id/reviews/:user_id", handlers.DeleteReview)

}
