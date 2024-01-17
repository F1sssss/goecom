package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// AddToCart

func AddToCart(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	user := c.Get("UserID")
	fmt.Println("User ID:", user)

	product := models.Product{}

	// Get the product ID from the URL
	productID := c.Param("id")

	quantity, err := strconv.Atoi(c.QueryParam("quantity"))

	// Check if the quantity is a number
	if err != nil {
		fmt.Println("Error converting quantity to int:", err)
		return c.JSON(http.StatusBadRequest, "Error converting quantity to int")
	}

	// Get the product from the database
	if err := db.Where("id = ?", productID).First(&product).Error; err != nil {
		fmt.Println("Error finding product:", err)
		return c.JSON(http.StatusInternalServerError, "Error finding product")
	}

	//Check if there is enough stock
	if product.Stock < int(quantity) {
		return c.JSON(http.StatusBadRequest, "Not enough stock")
	}

	// Get the user from the database
	userModel := models.User{}
	if err := db.Where("id = ?", user).First(&userModel).Error; err != nil {
		fmt.Println("Error finding user:", err)
		return c.JSON(http.StatusInternalServerError, "Error finding user")
	}

	ProductQuantity := models.ProductQuantity{Product: product, Quantity: int(quantity), ProductID: product.ID, ShoppingCartID: 0}

	ProductQuantity2 := db.Create(&ProductQuantity)

	fmt.Println("Product quantity:", ProductQuantity2)

	// Check if the user already has a shopping cart
	shoppingCart := models.ShoppingCart{}
	if err := db.Where("user_id = ?", user).First(&shoppingCart).Error; err != nil {

		// If the user doesn't have a shopping cart, create one
		shoppingCart = models.ShoppingCart{User: userModel, Products: []models.ProductQuantity{ProductQuantity}}

		if err := db.Create(&shoppingCart).Error; err != nil {
			fmt.Println("Error creating shopping cart:", err)
			return c.JSON(http.StatusInternalServerError, "Error creating shopping cart")
		}

		return c.JSON(http.StatusOK, shoppingCart)
	}

	shoppingCart.Products = append(shoppingCart.Products, ProductQuantity)

	// Check if the product is already in the shopping cart, in that case, just update the quantity
	for _, p := range shoppingCart.Products {

		fmt.Println("Product ID:", p.ProductID, "Product ID 2:", product.ID)

		if p.ProductID == product.ID {

			//Check if there is enough stock
			if product.Stock < int(quantity)+p.Quantity {
				return c.JSON(http.StatusBadRequest, "Not enough stock")
			}

			ProductQuantity.Quantity = quantity + p.Quantity

			fmt.Println("Product quantity:", ProductQuantity)

			if err := db.Save(&ProductQuantity).Error; err != nil {
				fmt.Println("Error updating shopping cart:", err)
				return c.JSON(http.StatusInternalServerError, "Error updating shopping cart")
			}
			return c.JSON(http.StatusOK, shoppingCart)
		}
	}

	ProductQuantity.ShoppingCartID = shoppingCart.ID

	// Add the product to the shopping cart
	shoppingCart.Products = append(shoppingCart.Products, ProductQuantity)

	fmt.Println("Shopping cart:", shoppingCart)

	// Update the shopping cart
	if err := db.Save(&shoppingCart).Error; err != nil {
		fmt.Println("Error updating shopping cart:", err)
		return c.JSON(http.StatusInternalServerError, "Error updating shopping cart")
	}

	return c.JSON(http.StatusOK, shoppingCart)

}
