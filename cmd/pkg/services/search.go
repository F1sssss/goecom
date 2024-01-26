package services

import (
	"net/http"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// SearchProducts godoc

func SearchProducts(c echo.Context) error {

	db := c.Get("db").(*gorm.DB)
	searchValue := c.QueryParam("search")

	products := []models.Product{}

	//Using Postgres full text search and pg_trgm extension (https://www.postgresql.org/docs/current/pgtrgm.html)
	//TODO: Maybe change to %>> or %<-> for better results
	if err := db.Where("description %> ?", searchValue).Or("name %> ?", searchValue).Limit(3).Find(&products).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}
