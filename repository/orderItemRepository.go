package repository

import "70_Off/entity"

type OrderItemRepository interface {
	SaveOrder(order *entity.OrderItems) (*entity.OrderItems, error)
	FindOrderById(id uint) (*entity.OrderItems, error)
	UpdateOrder(order *entity.OrderItems) (*entity.OrderItems, error)
	DeleteOrder(order *entity.OrderItems) error
	FindOrdersByUserId(userId uint) ([]entity.OrderItems, error)
	ListUserOrdersPaginated(userID uint, offset int, limit int) ([]entity.OrderItems, error)
}
