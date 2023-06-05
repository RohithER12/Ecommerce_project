package services

import (
	"70_Off/domain/entity"
	"70_Off/domain/repository"
)

type CartItemService interface {
	SaveCartItem(cartItem *entity.CartItem) (*entity.CartItem, error)
	FindCartItemByID(id uint) (*entity.CartItem, error)
	UpdateCartItem(cartItem *entity.CartItem) (*entity.CartItem, error)
	RemoveCartItem(cartItem *entity.CartItem) error
}

type cartItemService struct {
	cartItemRepo repository.CartItemRepository
}

func NewCartItemService(cartItemRepo repository.CartItemRepository) CartItemService {
	err := &cartItemService{cartItemRepo: cartItemRepo}
	if err != nil {
		return err
	}
	return nil
}

func (s *cartItemService) SaveCartItem(cartItem *entity.CartItem) (*entity.CartItem, error) {
	caritemUpdated, err := s.cartItemRepo.SaveCartItem(cartItem)
	if err != nil {
		return nil, err
	}
	return caritemUpdated, nil
}

func (s *cartItemService) FindCartItemByID(id uint) (*entity.CartItem, error) {
	cartitem, err := s.cartItemRepo.FindCartItemByID(id)
	if err != nil {
		return nil, err
	}
	return cartitem, nil
}

func (s *cartItemService) UpdateCartItem(cartItem *entity.CartItem) (*entity.CartItem, error) {
	cartitem, err := s.cartItemRepo.UpdateCartItem(cartItem)
	if err != nil {
		return nil, err
	}
	return cartitem, nil
}

func (s *cartItemService) RemoveCartItem(cartItem *entity.CartItem) error {

	err := s.cartItemRepo.DeleteCartItem(cartItem)
	if err != nil {
		return err
	}
	return nil
}
