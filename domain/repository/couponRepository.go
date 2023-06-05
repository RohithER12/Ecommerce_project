package repository

import "70_Off/domain/entity"

type CouponRepository interface {
	GetCouponByID(id uint) (*entity.Coupon, error)
	GetCouponByCouponCode(CouponCode string) (*entity.Coupon, error)
	CreateCoupon(coupon *entity.Coupon) error
	GetAllCoupons() ([]*entity.Coupon, error)
	UpdateCoupon(coupon *entity.Coupon) error
	DeleteCoupon(coupon *entity.Coupon) error
}
