package entity

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Name           string  `gorm:"not null"`
	OrderId        uint    `gorm:"not null"`
	TotalAmount    float64 `gorm:"not null"`
	DiscountAmount float64
	PayableAmount  float64
	PaymentID      uint `gorm:"not null"`
}

type Discount struct {
	gorm.Model
	Name   string
	Amount uint
	Limit  uint
}
