package models

import (
	"gorm.io/gorm"
)

// Role Model
type Role struct {
	gorm.Model
	RoleID       int    `json:"role_id"`
	RoleName     string `json:"role_name"`
	RoleSecLevel int    `json:"role_sec_level"`
}
