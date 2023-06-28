package repositoryImpl

import (
	"70_Off/entity"
	"errors"

	"gorm.io/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepoImpl(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}

func (r *CartRepo) SaveCart(cart *entity.Cart) (*entity.Cart, error) {
	if err := r.db.Create(&cart).Error; err != nil {
		return nil, errors.New("failed to savecart ")
	}
	return cart, nil
}

func (r *CartRepo) FindCartByID(id uint) (*entity.Cart, error) {
	var cart entity.Cart
	result := r.db.First(&cart, id)
	if result.Error != nil {
		return nil, errors.New("empty Cart")
	}
	return &cart, nil
}

func (r *CartRepo) UpdateCart(cart *entity.Cart) (*entity.Cart, error) {
	if err := r.db.Save(&cart).Error; err != nil {
		return nil, errors.New("failed to updateCart")
	}
	return cart, nil
}

func (r *CartRepo) DeleteCart(cart *entity.Cart) error {
	if err := r.db.Delete(&cart).Error; err != nil {
		return errors.New("failed to deleteCart ")
	}
	return nil
}

func (r *CartRepo) FindCartByUserID(userID uint) (*entity.Cart, error) {
	var cart entity.Cart
	if err := r.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, errors.New("failed to FindCartUserID")
	}
	return &cart, nil
}

func (r *CartRepo) FindCartsByUserID(userID uint) (*entity.Cart, error) {
	var cart entity.Cart
	if err := r.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}
