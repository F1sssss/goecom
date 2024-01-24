package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Add a review
func AddReview(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)
	userID := c.Get("UserID").(uint64)

	review := models.Review{}

	if err := c.Bind(&review); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	product_refer, err := strconv.Atoi(c.Param("id"))
	review.ProductRefer = uint(product_refer)
	review.UserRefer = uint(userID)

	if err != nil {
		fmt.Println("Error converting product_refer to int:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while converting product_refer to int")
	}

	// Check if the user already reviewed the product
	if err := db.Where("product_refer = ? AND user_refer = ?", review.ProductRefer, review.UserRefer).First(&review).Error; err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User already reviewed the product")
	}

	// Create the review
	if err := db.Create(&review).Error; err != nil {
		fmt.Println("Error creating review:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while creating review")
	}

	return c.JSON(http.StatusCreated, review)
}

// Get all reviews for a product
func GetProductReviews(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	reviews := []models.Review{}

	product_refer := c.Param("id")

	if err := db.Where("product_refer = ?", product_refer).Find(&reviews).Error; err != nil {
		fmt.Println("Error getting reviews:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while getting reviews")
	}

	return c.JSON(http.StatusOK, reviews)
}

// Update a review
func UpdateReview(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	review := models.Review{}

	if err := c.Bind(&review); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Check if the review exists
	if err := db.Where("id = ?", review.ID).First(&review).Error; err != nil {
		fmt.Println("Error getting review:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while getting review")
	}

	// Update the review
	if err := db.Save(&review).Error; err != nil {
		fmt.Println("Error updating review:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while updating review")
	}

	return c.JSON(http.StatusCreated, review)
}

// Delete a review
func DeleteReview(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	review := models.Review{}

	if err := c.Bind(&review); err != nil {
		fmt.Println("Error binding request body:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Delete the review
	if err := db.Delete(&review).Error; err != nil {
		fmt.Println("Error deleting review:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while deleting review")
	}

	return c.JSON(http.StatusCreated, review)
}

// Get all reviews for a user
func GetUserReviews(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	reviews := []models.Review{}

	user_refer := c.Param("user_refer")

	if err := db.Where("user_refer = ?", user_refer).Find(&reviews).Error; err != nil {
		fmt.Println("Error getting reviews:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error while getting reviews")
	}

	return c.JSON(http.StatusOK, reviews)
}
