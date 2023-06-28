package entity

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ProductDeatailsId uint
	Quantity          uint
}
