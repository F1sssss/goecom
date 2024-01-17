package handlers

import (
	"fmt"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Get All Factory
func GetAllFactory(c echo.Context, data interface{}) interface{} {
	db := c.Get("db").(*gorm.DB)
	result := db.Find(data)

	if result.Error != nil {
		fmt.Println("Error getting data from database:", result.Error)
		return result.Error
	}

	return nil
}

// Get One Factory
func GetOneFactory(c echo.Context, data interface{}) interface{} {
	db := c.Get("db").(*gorm.DB)
	result := db.First(data, c.Param("id"))

	if result.Error != nil {
		fmt.Println("Error getting data from database:", result.Error)
		return result.Error
	}

	return nil
}

// Create Factory
func CreateFactory(c echo.Context, data interface{}) interface{} {

	db := c.Get("db").(*gorm.DB)

	// Bind the JSON request body to the Model struct
	if err := c.Bind(data); err != nil {
		fmt.Println("Error binding request body:", err)
		return err
	}

	// Create the Model in the database
	result := db.Create(data)
	if result.Error != nil {
		fmt.Println("Error creating data:", result.Error)
		return result.Error
	}

	return nil
}

// Update Factory
func UpdateFactory(c echo.Context, data interface{}) interface{} {

	db := c.Get("db").(*gorm.DB)

	//Get the data from the database
	result := db.First(data, c.Param("id"))
	if result.Error != nil {
		fmt.Println("Error getting data from database:", result.Error)
		return result.Error
	}

	// Bind the JSON request body to the Model struct
	if err := c.Bind(data); err != nil {
		fmt.Println("Error binding request body:", err)
		return err
	}

	// Update the Model in the database
	result = db.Save(data)
	if result.Error != nil {
		fmt.Println("Error updating data:", result.Error)
		return result.Error
	}

	return nil
}

// Delete Factory
func DeleteFactory(c echo.Context, data interface{}) interface{} {

	db := c.Get("db").(*gorm.DB)

	//Get the data from the database
	result := db.First(data, c.Param("id"))
	if result.Error != nil {
		fmt.Println("Error getting data from database:", result.Error)
		return result.Error
	}

	// Delete the Model in the database
	result = db.Delete(data)
	if result.Error != nil {
		fmt.Println("Error deleting data:", result.Error)
		return result.Error
	}

	return nil
}
