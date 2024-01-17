package main

import (
	"fmt"
	"log"

	"github.com/F1sssss/goecom/cmd/pkg/database"
	"github.com/F1sssss/goecom/cmd/pkg/middleware"
	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/F1sssss/goecom/cmd/pkg/routes"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func InitializeDatabaseConnection() *gorm.DB { // Connect to the database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// Check if the connection is successful
	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		log.Fatalf("Error checking database connection: %v", err)
	}

	fmt.Printf("Connected to the database successfully! Result: %d\n", result)

	return db
}

func MigrateSchemas(e *gorm.DB) {

	if err := e.AutoMigrate(&models.ProductQuantity{}); err != nil {
		panic(err)
	}

	if err := e.AutoMigrate(&models.ShoppingCart{}); err != nil {
		panic(err)
	}

	if err := e.AutoMigrate(&models.Product{}); err != nil {
		panic(err)
	}

	if err := e.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	if err := e.AutoMigrate(&models.Role{}); err != nil {
		panic(err)
	}

	if err := e.AutoMigrate(&models.Review{}); err != nil {
		panic(err)
	}

}

func main() {

	db := InitializeDatabaseConnection()
	MigrateSchemas(db)

	e := echo.New()

	// Middleware
	e.Use(middleware.DatabaseMiddleware)

	// Routes
	routes.InitProductRoutes(e)
	routes.InitAuthRoutes(e)
	routes.InitUserRoutes(e)

	// Start server
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	fmt.Println("Server started successfully")

}
