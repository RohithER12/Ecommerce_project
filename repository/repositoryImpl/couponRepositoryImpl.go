package repositoryImpl

import (
	"70_Off/entity"
	"time"

	"gorm.io/gorm"
)

type CouponRepo struct {
	DB *gorm.DB
}

func NewCouponRepoImpl(db *gorm.DB) *CouponRepo {
	return &CouponRepo{DB: db}
}

func (r *CouponRepo) CreateCoupon(coupon *entity.Coupon) error {
	result := r.DB.Create(coupon)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CouponRepo) GetCouponByID(id uint) (*entity.Coupon, error) {
	var coupon entity.Coupon
	result := r.DB.First(&coupon, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if r.isCouponExpired(&coupon) {
		return nil, nil // or return an appropriate error indicating the coupon is expired
	}

	return &coupon, nil
}

func (r *CouponRepo) GetAllCoupons() ([]*entity.Coupon, error) {
	var coupons []*entity.Coupon
	result := r.DB.Find(&coupons)
	if result.Error != nil {
		return nil, result.Error
	}

	filteredCoupons := make([]*entity.Coupon, 0, len(coupons))
	for _, coupon := range coupons {
		if !r.isCouponExpired(coupon) {
			filteredCoupons = append(filteredCoupons, coupon)
		}
	}

	return filteredCoupons, nil
}

func (r *CouponRepo) UpdateCoupon(coupon *entity.Coupon) error {
	result := r.DB.Save(coupon)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CouponRepo) DeleteCoupon(coupon *entity.Coupon) error {
	result := r.DB.Delete(coupon)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CouponRepo) isCouponExpired(coupon *entity.Coupon) bool {
	return coupon.Expiration_date.Before(time.Now())
}

func (r *CouponRepo) GetCouponByCouponCode(couponCode string) (*entity.Coupon, error) {
	var coupon entity.Coupon
	result := r.DB.Where("coupon_code = ?", couponCode).First(&coupon)
	if result.Error != nil {
		return nil, result.Error
	}

	if r.isCouponExpired(&coupon) {
		return nil, nil
	}

	return &coupon, nil
}
