package services

// import (
// 	"70_Off/domain/entity"
// 	"70_Off/domain/repository"
// 	"fmt"
// )

// type CartServiceImpl struct {
// 	cartRepo repository.CartRepository
// }

// type CartService interface {
// 	SaveCart(cart *entity.Cart) (*entity.Cart, error)
// 	FindCartByID(id uint) (*entity.Cart, error)
// 	UpdateCart(cart *entity.Cart) (*entity.Cart, error)
// 	DeleteCart(cart *entity.Cart) error
// 	FindCartsByUserID(userID uint) ([]entity.Cart, error)
// }

// func NewCartServiceImpl(cartRepo repository.CartRepository) *CartServiceImpl {
// 	return &CartServiceImpl{cartRepo: cartRepo}
// }

// func (s *CartServiceImpl) SaveCart(cart *entity.Cart) (*entity.Cart, error) {
// 	return s.cartRepo.SaveCart(cart)
// }

// func (s *CartServiceImpl) FindCartByID(id uint) (*entity.Cart, error) {
// 	cart, err := s.cartRepo.FindCartByID(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to find cart: %w", err)
// 	}
// 	return cart, nil
// }

// func (s *CartServiceImpl) UpdateCart(cart *entity.Cart) (*entity.Cart, error) {
// 	updatedCart, err := s.cartRepo.UpdateCart(cart)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update cart: %w", err)
// 	}
// 	return updatedCart, nil
// }

// func (s *CartServiceImpl) DeleteCart(cart *entity.Cart) error {
// 	err := s.cartRepo.DeleteCart(cart)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete cart: %w", err)
// 	}
// 	return nil
// }

// func (s *CartServiceImpl) FindCartsByUserID(userID uint) ([]entity.Cart, error) {
// 	carts, err := s.cartRepo.FindCartsByUserID(userID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to find carts by user ID: %w", err)
// 	}
// 	return carts, nil
// }
