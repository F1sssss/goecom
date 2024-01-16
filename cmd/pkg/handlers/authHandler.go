package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/F1sssss/goecom/cmd/pkg/utils"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Register user
func Register(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	user := models.User{}

	if err := c.Bind(&user); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Check if the user already exists
	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User already exists")
	}

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while hashing password")
	}

	user.Password = hashedPassword
	user.Verified = false

	// Create the user
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Error creating user:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while creating user")
	}

	return c.JSON(http.StatusCreated, user)
}

// Login user
func Login(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	user := models.User{}

	if err := c.Bind(&user); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Check if the user exists
	if err := db.Where("username = ?", user.Username).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	// Compare the passwords
	if utils.ComparePasswords(user.Password, user.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Username, user.Email, user.Role)

	if err != nil {
		fmt.Println("Error generating JWT token:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while generating JWT token")
	}

	// Create and set a cookie
	c.SetCookie(createCookie("token", token, time.Now().Add(24*time.Hour)))

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

}

func createCookie(name string, value string, expires time.Time) *http.Cookie {
	return &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expires,
	}
}
