package entity

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	UserID          uint
	AddressID       uint
	TotalBillAmount float64
	PaymentTypeID   uint
}

type OrderItems struct {
	gorm.Model
	OrderID          uint
	ProductDetailsID uint
	Quantity         uint
	Price            float64
	Status           string `gorm:"not null;default:'order placed'"`
}
