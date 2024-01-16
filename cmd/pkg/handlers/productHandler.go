package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var lock = sync.Mutex{}

// GetProducts returns all products
func GetProducts(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	// Retrieve products from the database
	var products []models.Product
	result := db.Find(&products)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, products)

}

// GetProduct returns a single product
func GetProduct(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	// Retrieve product from the database
	var product models.Product
	result := db.First(&product, c.Param("id"))
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, product)

}

// CreateProduct creates a new product
func CreateProduct(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	db := c.Get("db").(*gorm.DB)

	// Bind the JSON request body to the Product struct
	var product models.Product
	if err := c.Bind(&product); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Create the product in the database
	result := db.Create(&product)
	if result.Error != nil {
		fmt.Println("Error creating product:", result.Error)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Return the created product in the response
	return c.JSON(http.StatusCreated, product)
}

// UpdateProduct updates an existing product
func UpdateProduct(c echo.Context) error {

	lock.Lock()
	defer lock.Unlock()

	db := c.Get("db").(*gorm.DB)

	// Retrieve product from the database
	var product models.Product
	result := db.First(&product, c.Param("id"))
	if result.Error != nil {
		return result.Error
	}

	// Bind the JSON request body to the Product struct
	if err := c.Bind(&product); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Update the product in the database
	result = db.Save(&product)
	if result.Error != nil {
		fmt.Println("Error updating product:", result.Error)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Return the updated product in the response
	return c.JSON(http.StatusOK, product)

}

func DeleteProduct(c echo.Context) error {

	lock.Lock()
	defer lock.Unlock()

	db := c.Get("db").(*gorm.DB)

	// Retrieve product from the database
	var product models.Product
	result := db.First(&product, c.Param("id"))
	if result.Error != nil {
		return result.Error
	}

	// Delete the product in the database
	result = db.Delete(&product)
	if result.Error != nil {
		fmt.Println("Error deleting product:", result.Error)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Return the deleted product in the response
	return c.String(http.StatusOK, "Product deleted successfully")

}
