package repository

import "70_Off/entity"

type CartRepository interface {
	SaveCart(cart *entity.Cart) (*entity.Cart, error)
	FindCartByID(id uint) (*entity.Cart, error)
	UpdateCart(cart *entity.Cart) (*entity.Cart, error)
	DeleteCart(cart *entity.Cart) error
	FindCartsByUserID(userID uint) (*entity.Cart, error)
	FindCartByUserID(userID uint) (*entity.Cart, error)
}
