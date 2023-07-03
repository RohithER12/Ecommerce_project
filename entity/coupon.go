package entity

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	CouponCode            string
	DiscountPercentage    float64
	MaximumDiscountAmount float64
	Expiration_date       time.Time
}
