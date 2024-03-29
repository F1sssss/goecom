package middleware

import (
	"fmt"

	"github.com/F1sssss/goecom/cmd/pkg/database"
	"github.com/labstack/echo"
)

// DatabaseMiddleware creates a database connection and injects it into the context
func DatabaseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Borrow a connection from the connection pool
		conn := database.GetDB()

		if conn == nil {
			fmt.Println("Error getting database connection")
			return echo.ErrInternalServerError
		}

		// Inject the connection into the context
		c.Set("db", conn)
		// Call the next handler
		err := next(c)

		return err
	}
}
