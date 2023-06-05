package entity

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentTypeId uint    `gorm:"not null"`
	PaymentStatus bool    `gorm:"not null;default:false"`
	PayableAmount float64 `gorm:"not null"`
	OrderId       uint    `gorm:"not null"`
}

type PaymentType struct {
	gorm.Model
	PaymentType string
}
