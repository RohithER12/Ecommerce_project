package deliverymodels

import "time"

type AdminSignupInputs struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber" gorm:"unique_index"`
	Password    string `json:"password"`
}

type AdmionLoginInputs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddProductInputs struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required" `
	BrandId     uint   `json:"branid" `
}

type AddProductsDeatils struct {
	ProductId uint    `json:"productid" `
	SizeId    uint    `json:"sizeid" `
	Quantity  uint    `json:"quantity" binding:"required" `
	Price     float64 `json:"price" `
}

type UpdateProductInput struct {
	ProductId uint `json:"productid" `
	// AdminId     uint   `json:"adminid" `
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required" `
	BrandId     uint   `json:"branid" `
}

type UpdateProductsDeatils struct {
	ProductItemId uint `json:"productItemId" `
	// AdminId       uint `json:"adminId" `
	ProductId uint `json:"productId" `
	SizeId    uint `json:"sizeid" `
	Quantity  uint `json:"quantity" binding:"required" `
	Price     float64
}

type OrderReport struct {
	NoOfOrders         uint
	TotalAmount        float64
	MonthlyOrders      uint
	MonthlyTotalAmount float64
	YearlyOrders       uint
	YearlyTotalAmount  float64
	WeeklyOrders       uint
	WeeklyTotalAmount  float64
}

type Coupon struct {
	CouponCode            string    `json:"couponCode" `
	DiscountPercentage    uint      `json:"discount" `
	MaximumDiscountAmount uint      `json:"maximumdiscount" `
	Expiration_date       time.Time `json:"expirDate" `
}
