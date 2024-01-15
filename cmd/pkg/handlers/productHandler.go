package handlers

import (
	"github.com/F1sssss/goecom/cmd/pkg/database"
	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
)

// GetProducts returns all products
func GetProducts(c echo.Context) error {

	// Connect to the database
	db, err := database.Connect()

	if err != nil {
		return err
	}

	// Retrieve products from the database
	var products []models.Product
	result := db.Find(&products)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(200, products)

}
