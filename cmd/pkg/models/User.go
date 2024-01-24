package models

import (
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Email              string   `json:"email" gorm:"unique"`
	Username           string   `json:"username" gorm:"unique"`
	Password           string   `json:"password"`
	Name               string   `json:"name"`
	LastName           string   `json:"last_name"`
	DateOfBirth        string   `json:"date_of_birth"`
	Reviews            []Review `gorm:"foreignKey:UserRefer"`
	Role               Role     `json:"role" gorm:"foreignKey:RoleID"`
	Verified           bool     `json:"verified"`
	ConfirmationToken  string   `json:"confirmation_token"`
	ResetPasswordToken string   `json:"reset_password_token"`
	RecoveryTokenTime  string   `json:"recovery_token_time"`
}
