package repository

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/entity"
)

type CartItemRepository interface {
	SaveCartItem(cartItem *entity.CartItem) (*entity.CartItem, error)
	FindCartItemByID(id uint) (*entity.CartItem, error)
	UpdateCartItem(cartItem *entity.CartItem) (*entity.CartItem, error)
	DeleteCartItem(cartItem *entity.CartItem) error
	FindCartItemsByCartID(cartID uint) (*entity.CartItem, error)
	FindCartItemByCartAndProduct(cartID uint, productID uint) (*entity.CartItem, error)
	FindAllCartItemsByCartID(cartID uint) ([]entity.CartItem, error)
	DeleteCartItemWithQuantity(cartItemID, quantity uint) error
	FindCartItemByUserIdAndProductDetailsID(userId, productDetailsId uint) (*entity.CartItem, error)
	FindAllCartItemsByCartIDForUser(cartID uint) ([]deliverymodels.ViewCartDetail, error)
}
