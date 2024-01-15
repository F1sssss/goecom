package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/F1sssss/goecom/cmd/pkg/database"
	"github.com/F1sssss/goecom/cmd/pkg/handlers"
	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}

	// Check if the connection is successful
	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		log.Fatalf("Error checking database connection: %v", err)
	}

	fmt.Printf("Connected to the database successfully! Result: %d\n", result)

	Product := models.Product{}
	db.AutoMigrate(&Product)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/products", handlers.GetProducts)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	fmt.Println("Server started successfully")
	fmt.Println(result)

}
