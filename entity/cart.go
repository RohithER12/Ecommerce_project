package entity

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID           uint `json:"cartId"`
	ProductDetailsID uint `json:"productId"`
	SizeID           uint
	Quantity         uint `json:"quantity"`
	Price            float64
}

type Cart struct {
	gorm.Model
	UserID          uint `json:"userId"`
	TotalCartAmount float64
}
