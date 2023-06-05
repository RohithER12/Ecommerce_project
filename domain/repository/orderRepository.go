package repository

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/domain/entity"
	"time"
)

type OrderRepository interface {
	SaveOrder(order *entity.Orders) (*entity.Orders, error)
	FindOrderById(id uint) (*entity.Orders, error)
	UpdateOrder(order *entity.Orders) (*entity.Orders, error)
	DeleteOrder(order *entity.Orders) error
	FindOrdersByUserId(userId uint) ([]entity.Orders, error)
	ListUserOrdersPaginated(userID uint, offset int, limit int) ([]entity.Orders, error)
	SaveOrderItem(order *entity.OrderItems) (*entity.OrderItems, error)
	FindOrderItemById(id uint) (*entity.OrderItems, error)
	UpdateOrderItem(order *entity.OrderItems) (*entity.OrderItems, error)
	DeleteOrderItem(order *entity.OrderItems) error
	FindOrderItemByUserId(userId uint) ([]entity.OrderItems, error)
	ListUserOrderItemPaginated(userID uint, offset int, limit int) ([]deliverymodels.ViewOrdersForUsers, error)
	GetOrderStatistics(startDate, endDate time.Time) (int64, float64, error)
	FindAllOrderItemsByOrderID(orderID uint) ([]deliverymodels.ViewOrdersForUsers, error)
}
