package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" binding:"required" gorm:"not null" gorm:"unique_index"`
	Description string `json:"description" binding:"required" gorm:"not null"`
	Deleted     bool   `gorm:"default:false"`
	AdminID     uint   `json:"-"`
	BrandID     uint   `json:"brandId"`
	Slug        string ` gorm:"unique_index"` //gorm:"not null"
}

type ProductDetails struct {
	gorm.Model
	AdminID   uint
	Price     float64 `json:"price" binding:"required" gorm:"not null"`
	SizeID    uint    `json:"size_id" binding:"required" gorm:"not null"`
	ProductID uint    `json:"product_id" binding:"required" gorm:"not null"`
	Quantity  uint
	Slug      string ` gorm:"unique_index"` //gorm:"not null"
}
