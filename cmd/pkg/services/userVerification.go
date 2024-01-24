package services

import (
	"fmt"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"gorm.io/gorm"

	"github.com/labstack/echo"
)

// Verify verifies the user's email address

func VerifyUser(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	token := c.QueryParam("token")
	id := c.QueryParam("id")

	user := models.User{}

	fmt.Println(token, id)

	if err := db.Where("id = ?", id).Where("confirmation_token = ?", token).First(&user).Error; err != nil {
		fmt.Println("Error finding user:", err)
		return echo.NewHTTPError(404, "User not found or token is invalid")
	}

	if err := db.Model(&models.User{}).Where("id = ?", id).Where("confirmation_token = ?", token).Update("verified", true).Error; err != nil {
		return echo.NewHTTPError(500, "Internal Server Error while updating user")
	}

	return c.JSON(200, "User verified")

}
