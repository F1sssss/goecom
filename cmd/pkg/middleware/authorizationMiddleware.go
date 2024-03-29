package middleware

import (
	"fmt"
	"net/http"

	"github.com/F1sssss/goecom/cmd/pkg/utils"
	"github.com/labstack/echo"
)

func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("token")
		//Check if the cookie is set
		if err != nil {
			fmt.Println("Error getting cookie:", err)
			return c.JSON(http.StatusForbidden, "Unauthorized")
		}

		// Check if the token is valid
		token, claims, err := utils.ParseJWT(cookie.Value)

		if err != nil {
			fmt.Println("Error parsing JWT:", err)
			return c.JSON(http.StatusForbidden, "Unauthorized")
		}

		if !token.Valid {
			fmt.Println("Invalid token")
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("Role", claims.Role.RoleName)
		c.Set("UserID", claims.UserID)

		return next(c)

	}
}

func HasRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("Role")

			if claims != role {
				return c.JSON(http.StatusForbidden, "Unauthorized")
			}

			return next(c)

		}
	}
}

// IsMe checks if the user is the same as the one in the token, it is used for updating and deleting own resources only
// It is used by puttin user_id in the path and checking if it is the same as the one in the token
func IsMe(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("UserID")

		if claims != c.Param("user_id") {
			return c.JSON(http.StatusForbidden, "Unauthorized")
		}

		return next(c)

	}
}
