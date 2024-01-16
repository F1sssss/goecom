package models

import (
	"gorm.io/gorm"
)

// Product model

type Product struct {
	gorm.Model
	SKU         string  `gorm:"unique;not null" json:"sku"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `gorm:"not null" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"not null" json:"stock"`
	CategoryID  int     `gorm:"not null" json:"category_id"`
}
