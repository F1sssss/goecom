package handlers

import (
	"fmt"
	"net/http"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
)

// GetUsers returns all users
func GetUsers(c echo.Context) error {

	var users []models.User

	if err := GetAllFactory(c, &users); err != nil {
		fmt.Println("Error getting users:", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, users)

}

// GetUser returns a single user
func GetUser(c echo.Context) error {

	var user models.User

	if err := GetOneFactory(c, &user); err != nil {
		fmt.Println("Error getting user:", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, user)

}

// UpdateUser updates a user
func UpdateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var user models.User

	if err := UpdateFactory(c, &user); err != nil {
		fmt.Println("Error updating user:", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, user)

}

// DeleteUser deletes a user
func DeleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var user models.User

	if err := DeleteFactory(c, &user); err != nil {
		fmt.Println("Error deleting user:", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusNoContent, "User deleted")
}
