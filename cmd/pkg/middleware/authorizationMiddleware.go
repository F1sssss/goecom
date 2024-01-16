package middleware

import (
	"fmt"

	"github.com/F1sssss/goecom/cmd/pkg/utils"
	"github.com/labstack/echo"
)

func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("token")

		if err != nil {
			fmt.Println("Error getting cookie:", err)
			return c.JSON(401, "Unauthorized")
		}

		// Check if the token is valid
		token, claims, err := utils.ParseJWT(cookie.Value)

		if err != nil {
			fmt.Println("Error parsing JWT:", err)
			return c.JSON(401, "Unauthorized")
		}

		if !token.Valid {
			fmt.Println("Invalid token")
			return c.JSON(401, "Unauthorized")
		}

		fmt.Println("JWT claims:", claims.Role.RoleName)

		return next(c)

	}
}
