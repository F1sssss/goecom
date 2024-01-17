package models

import (
	"gorm.io/gorm"
)

type ProductQuantity struct {
	gorm.Model
	ProductID      uint    `json:"product_id"`
	Product        Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity       int     `json:"quantity"`
	ShoppingCartID uint    `json:"shopping_cart_id"`
}

// ShoppingCart Model
type ShoppingCart struct {
	gorm.Model
	Products   []ProductQuantity `json:"products" gorm:"foreignKey:ShoppingCartID"`
	User       User              `json:"user" gorm:"foreignKey:UserID"`
	PriceTotal float64           `json:"price_total"`
	UserID     int64             `json:"user_id"`
}
