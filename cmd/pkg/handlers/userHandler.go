package handlers

import (
	"fmt"
	"net/http"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// GetUsers returns all users
func GetUsers(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	// Retrieve users from the database
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, users)

}

// Create User creates a new user
func CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	db := c.Get("db").(*gorm.DB)

	// Bind the JSON request body to the User struct
	var user models.User
	if err := c.Bind(&user); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Create a new user in the database
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, user)

}

// GetUser returns a single user

func GetUser(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	// Retrieve user from the database
	var user models.User
	result := db.First(&user, c.Param("id"))
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, user)

}

// UpdateUser updates a user

func UpdateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	db := c.Get("db").(*gorm.DB)

	// Retrieve user from the database
	var user models.User
	result := db.First(&user, c.Param("id"))
	if result.Error != nil {
		return result.Error
	}

	// Bind the JSON request body to the User struct
	if err := c.Bind(&user); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Save the updated user in the database
	result = db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user

func DeleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	db := c.Get("db").(*gorm.DB)

	// Retrieve user from the database
	var user models.User
	result := db.First(&user, c.Param("id"))
	if result.Error != nil {
		return result.Error
	}

	// Delete the user from the database
	result = db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusNoContent)
}
