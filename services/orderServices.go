package services

// import (
// 	"70_Off/domain/entity"
// 	"70_Off/domain/repository"
// )

// type OrderService interface {
// 	CreateOrder(userID uint, cartItemID uint, addressID uint) (*entity.Order, error)
// 	OrderProcessing(userID uint, cartId uint) (*entity.CartItem, error)
// 	// GetOrderByID(orderID uint) (*entity.Order, error)
// 	// UpdateOrder(order *entity.Order) (*entity.Order, error)
// 	// DeleteOrder(order *entity.Order) error
// 	// GetOrdersByUserID(userID uint) ([]entity.Order, error)
// }

// type OrderServiceImpl struct {
// 	orderRepo repository.OrderRepository
// }

// func NewOrderService(orderRepo repository.OrderRepository) *OrderServiceImpl {
// 	return &OrderServiceImpl{
// 		orderRepo: orderRepo,
// 	}
// }

// func (s *OrderServiceImpl) OrderProcessing(userID uint, cartId uint) (*entity.CartItem, error) {

// }

// func (s *OrderServiceImpl) CreateOrder(userID uint, cartItemID uint, addressID uint) (*entity.Order, error) {

// }

// func (s *OrderServiceImpl) GetOrderByID(orderID uint) (*entity.Order, error) {
// 	// Retrieve the order entity by ID using the order repository
// 	// Return the order entity or an error
// }

// func (s *OrderServiceImpl) UpdateOrder(order *entity.Order) (*entity.Order, error) {
// 	// Update the order entity using the order repository
// 	// Return the updated order entity or an error
// }

// func (s *OrderServiceImpl) DeleteOrder(order *entity.Order) error {
// 	// Delete the order entity using the order repository
// 	// Return an error if any
// }

// func (s *OrderServiceImpl) GetOrdersByUserID(userID uint) ([]entity.Order, error) {
// 	// Retrieve the orders for a specific user ID using the order repository
// 	// Return the list of orders or an error
// }
