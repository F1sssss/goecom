package models

import (
	"gorm.io/gorm"
)

// Review Model
type Review struct {
	gorm.Model
	ProductRefer uint   `gorm:"not null" json:"product_refer"`
	UserRefer    uint   `gorm:"not null" json:"user_refer"`
	Comment      string `json:"comment"`
	Rating       int    `json:"rating"`
}
