package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
)

// Lock for thread safety
var lock = sync.Mutex{}

// GetProducts returns all products
func GetProducts(c echo.Context) error {

	var products []models.Product

	if err := GetAllFactory(c, &products); err != nil {
		fmt.Println("Error getting products:", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, products)
}

// GetProduct returns a single product
func GetProduct(c echo.Context) error {

	var product models.Product

	if err := GetOneFactory(c, &product); err != nil {
		fmt.Println("Error getting product:", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product
func CreateProduct(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var product models.Product

	if err := CreateFactory(c, &product); err != nil {
		fmt.Println("Error creating product:", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusCreated, product)

}

// UpdateProduct updates an existing product
func UpdateProduct(c echo.Context) error {

	lock.Lock()
	defer lock.Unlock()

	var product models.Product

	if err := GetOneFactory(c, &product); err != nil {
		fmt.Println("Error getting product:", err)
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)

}

// DeleteProduct deletes an existing product
func DeleteProduct(c echo.Context) error {

	lock.Lock()
	defer lock.Unlock()

	if err := DeleteFactory(c, &models.Product{}); err != nil {
		fmt.Println("Error deleting product:", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, "Product deleted")

}
