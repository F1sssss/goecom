package models

import (
	"gorm.io/gorm"
)

// Review Model
type Review struct {
	gorm.Model
	Product Product `json:"product_id" gorm:"foreignKey:id"`
	User    User    `json:"user_id" gorm:"foreignKey:id"`
	Comment string  `json:"comment"`
	Rating  int     `json:"rating"`
}
