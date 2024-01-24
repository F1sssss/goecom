package routes

import (
	"github.com/F1sssss/goecom/cmd/pkg/handlers"
	"github.com/F1sssss/goecom/cmd/pkg/services"
	"github.com/labstack/echo"
)

// Init initializes all routes
func InitAuthRoutes(e *echo.Echo) {

	e.GET("/login", handlers.Login)
	e.POST("/register", handlers.Register)
	//e.GET("/logout", handlers.Logout)
	//e.POST("/forgot-password", handlers.ForgotPassword)
	//e.POST("/reset-password", handlers.ResetPassword)
	//e.POST("/change-password", handlers.ChangePassword)
	e.GET("/verify/", services.VerifyUser)
	//e.POST("/resend-verification-email", handlers.ResendVerificationEmail)
	//e.GET("/me", handlers.Me)
}
