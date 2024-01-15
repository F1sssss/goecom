package models

import (
	"gorm.io/gorm"
)

// Product model

type Product struct {
	gorm.Model
	SKU         string  `gorm:"unique;not null"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	CategoryID  int     `gorm:"not null"`
}
