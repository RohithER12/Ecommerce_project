package repositoryImpl

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/entity"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepoImpl(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) SaveOrder(order *entity.Orders) (*entity.Orders, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepo) FindOrderById(id uint) (*entity.Orders, error) {
	var order entity.Orders
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepo) UpdateOrder(order *entity.Orders) (*entity.Orders, error) {
	if err := r.db.Save(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepo) DeleteOrder(order *entity.Orders) error {
	if err := r.db.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepo) FindOrdersByUserId(userId uint) ([]entity.Orders, error) {
	var orders []entity.Orders
	if err := r.db.Where("user_id = ?", userId).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepo) ListUserOrdersPaginated(userID uint, offset int, limit int) ([]entity.Orders, error) {

	var orders []entity.Orders

	err := r.db.Table("orders").
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepo) SaveOrderItem(order *entity.OrderItems) (*entity.OrderItems, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepo) FindOrderItemById(id uint) (*entity.OrderItems, error) {
	var orderItem entity.OrderItems
	if err := r.db.First(&orderItem, id).Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (r *OrderRepo) UpdateOrderItem(order *entity.OrderItems) (*entity.OrderItems, error) {
	if err := r.db.Save(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepo) DeleteOrderItem(order *entity.OrderItems) error {
	if err := r.db.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepo) FindOrderItemByUserId(userId uint) ([]entity.OrderItems, error) {
	var orderItems []entity.OrderItems
	if err := r.db.Where("user_id = ?", userId).Find(&orderItems).Error; err != nil {
		return nil, err
	}
	return orderItems, nil
}
func (r *OrderRepo) ListUserOrderItemPaginated(userID uint, offset int, limit int) ([]deliverymodels.ViewOrdersForUsers, error) {
	var orderItems []entity.OrderItems
	var viewOrders []deliverymodels.ViewOrdersForUsers
	var orders []entity.Orders

	// Retrieve orders for the specified user with pagination
	if err := r.db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, err
	}

	for _, order := range orders {
		// Retrieve order items for each order
		if err := r.db.Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
			return nil, errors.New("no data found")
		}

		for _, item := range orderItems {
			productDetails := entity.ProductDetails{}
			if err := r.db.Where("id = ?", item.ProductDetailsID).First(&productDetails).Error; err != nil {
				return nil, errors.New("no data found")
			}

			product := entity.Product{}
			if err := r.db.Where("id = ?", productDetails.ProductID).First(&product).Error; err != nil {
				return nil, errors.New("no data found")
			}

			viewOrder := deliverymodels.ViewOrdersForUsers{
				OrderID:  item.OrderID,
				Product:  product.Name,
				Quantity: item.Quantity,
				Price:    productDetails.Price,
				Status:   item.Status,
			}

			viewOrders = append(viewOrders, viewOrder)
		}
	}

	return viewOrders, nil
}

func (r *OrderRepo) FindAllOrderItemsByOrderID(orderID uint) ([]deliverymodels.ViewOrdersForUsers, error) {
	var orderItems []entity.OrderItems
	var viewOrders []deliverymodels.ViewOrdersForUsers

	if err := r.db.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		return nil, errors.New("no data found")
	}

	for _, item := range orderItems {
		productDetails := entity.ProductDetails{}
		if err := r.db.Where("id = ?", item.ProductDetailsID).First(&productDetails).Error; err != nil {
			return nil, errors.New("no data found")
		}

		product := entity.Product{}
		if err := r.db.Where("id = ?", productDetails.ProductID).First(&product).Error; err != nil {
			return nil, errors.New("no data found")
		}

		viewOrder := deliverymodels.ViewOrdersForUsers{
			OrderID:  item.OrderID,
			Product:  product.Name,
			Quantity: item.Quantity,
			Price:    productDetails.Price,
			Status:   item.Status,
		}

		viewOrders = append(viewOrders, viewOrder)
	}

	return viewOrders, nil
}

func (r *OrderRepo) GetOrderStatistics(startDate, endDate time.Time) (int64, float64, error) {
	var orderCount int64
	var totalAmount float64

	type Result struct {
		OrderCount  int64
		TotalAmount float64
	}
	var result Result

	query := r.db.Table("orders").
		Where("created_at >= ? AND created_at <= ?", startDate, endDate).
		Select("count(*) as order_count, coalesce(sum(total_bill_amount), 0) as total_amount").
		Scan(&result)

	if query.Error != nil {
		fmt.Println("Failed to retrieve orders:", query.Error)
		return 0, 0, query.Error
	}

	orderCount = result.OrderCount
	totalAmount = result.TotalAmount

	return orderCount, totalAmount, nil
}

// func (r *OrderRepo) CheckOutCart(addressID, userID uint) (*entity.Orders, *entity.Invoice, error) {
// 	tx := r.db.Begin() // Start a database transaction

// 	// Rollback the transaction if an error occurs
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	address, err := r.userRepo.GetAddressById(addressID)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, nil, err
// 	}

// 	cart, err := r.cartRepo.FindCartByUserID(userID)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, nil, err
// 	}

// 	cartItems, err := r.cartItemRepo.FindAllCartItemsByCartID(cart.ID)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, nil, err
// 	}

// 	totalAmount := cart.TotalCartAmount

// 	order := &entity.Orders{
// 		UserID:          userID,
// 		AddressID:       addressID,
// 		TotalBillAmount: totalAmount,
// 	}

// 	createdOrder, err := r.SaveOrder(order)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, nil, err
// 	}

// 	for _, cartItem := range cartItems {
// 		orderItem := &entity.OrderItems{
// 			OrderID:          createdOrder.ID,
// 			ProductDetailsID: cartItem.ProductDetailsID,
// 			Quantity:         cartItem.Quantity,
// 			Price:            cartItem.Price,
// 			Status:           "order placed",
// 		}

// 		_, err = r.SaveOrderItem(orderItem)
// 		if err != nil {
// 			tx.Rollback()
// 			return nil, nil, err
// 		}

// 		err = r.cartItemRepo.DeleteCartItem(&cartItem)
// 		if err != nil {
// 			tx.Rollback()
// 			return nil, nil, err
// 		}
// 	}

// 	payment := entity.Payment{
// 		OrderID:       createdOrder.ID,
// 		PaymentTypeID: 1, // Assuming the payment type ID is 1 for Cash on Delivery
// 		PayableAmount: totalAmount,
// 	}

// 	_, err = r.paymentRepo.SavePayment(&payment)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, nil, err
// 	}

// 	invoice := &entity.Invoice{
// 		Name:        address.User.FirstName + " " + address.User.LastName,
// 		OrderID:     createdOrder.ID,
// 		TotalAmount: totalAmount,
// 		BillAmount:  totalAmount,
// 		PaymentID:   payment.ID,
// 	}

// 	invoiceCreated, err := r.invoiceRepo.SaveInvoice(invoice)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, nil, err
// 	}

// 	tx.Commit() // Commit the transaction

// 	return createdOrder, invoiceCreated, nil
// }
