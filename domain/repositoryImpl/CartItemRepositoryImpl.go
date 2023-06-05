package repositoryImpl

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/domain/entity"
	"errors"
	"log"

	"gorm.io/gorm"
)

type CartItemRepo struct {
	db *gorm.DB
}

func NewCartItemRepoImpl(db *gorm.DB) *CartItemRepo {
	return &CartItemRepo{db: db}
}

func (r *CartItemRepo) SaveCartItem(cartItem *entity.CartItem) (*entity.CartItem, error) {
	tx := r.db.Begin()
	if err := tx.Create(&cartItem).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to save cartitem")
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to commit transaction")
	}
	return cartItem, nil
}

func (r *CartItemRepo) FindCartItemByID(id uint) (*entity.CartItem, error) {
	var cartItem entity.CartItem
	if err := r.db.First(&cartItem, id).Error; err != nil {
		return nil, errors.New("failed to FindCartItemByID cartitem")
	}
	return &cartItem, nil
}

func (r *CartItemRepo) UpdateCartItem(cartItem *entity.CartItem) (*entity.CartItem, error) {
	err := r.db.Model(&entity.CartItem{}).Where("id = ?", cartItem.ID).Save(cartItem).Error
	if err != nil {
		log.Println("Failed to update cart item:", err)
		return nil, errors.New("failed to update cart item")
	}
	return cartItem, nil
}

func (r *CartItemRepo) DeleteCartItem(cartItem *entity.CartItem) error {
	if err := r.db.Delete(&cartItem).Error; err != nil {
		return errors.New("failed to deleteCartItem ")
	}
	return nil
}

func (r *CartItemRepo) FindCartItemsByCartID(cartID uint) (*entity.CartItem, error) {
	var cartItems entity.CartItem
	if err := r.db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		return nil, errors.New("failed to FindCartItemByCartID ")
	}
	return &cartItems, nil
}

func (r *CartItemRepo) FindAllCartItemsByCartID(cartID uint) ([]entity.CartItem, error) {
	var cartItems []entity.CartItem
	if err := r.db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no cart items found for the given cart ID")
		}
		return nil, err
	}
	return cartItems, nil
}

func (r *CartItemRepo) FindCartItemByUserIdAndProductDetailsID(userId, productDetailsId uint) (*entity.CartItem, error) {
	cart := &entity.Cart{}
	if err := r.db.Where("user_id = ?", userId).First(cart).Error; err != nil {
		return nil, errors.New("failed to find cart for user")
	}

	cartItem := &entity.CartItem{}
	if err := r.db.Where("cart_id = ? AND product_details_id = ?", cart.ID, productDetailsId).First(cartItem).Error; err != nil {
		return nil, errors.New("failed to find cart item by user ID and product details ID")
	}

	return cartItem, nil
}

func (r *CartItemRepo) FindCartItemByCartAndProduct(cartID uint, productID uint) (*entity.CartItem, error) {
	var cartItem entity.CartItem
	if err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartItem).Error; err != nil {
		return nil, errors.New("failed to FindCartItemByCartAndProduct")
	}
	return &cartItem, nil
}

func (r *CartItemRepo) DeleteCartItemWithQuantity(cartItemID, quantity uint) error {
	var cartItem entity.CartItem
	if err := r.db.First(&cartItem, cartItemID).Error; err != nil {
		return errors.New("failed to find cart item by ID")
	}
	cartItem.Price = (cartItem.Price / float64(cartItem.Quantity)) * (float64(cartItem.Quantity - quantity))
	cartItem.Quantity -= quantity

	if cartItem.Quantity <= 0 {
		if err := r.db.Delete(&cartItem).Error; err != nil {
			return errors.New("failed to delete cart item")
		}
	} else {
		if err := r.db.Save(&cartItem).Error; err != nil {
			return errors.New("failed to update cart item")
		}
	}

	return nil
}

func (r *CartItemRepo) FindAllCartItemsByCartIDForUser(cartID uint) ([]deliverymodels.ViewCartDetail, error) {
	var viewCartDetails []deliverymodels.ViewCartDetail

	// Perform the join operations to fetch the required data
	err := r.db.
		Table("cart_items").
		Select("products.name, product_details.price, products.description, sizes.size, brands.brand, cart_items.quantity").
		Joins("JOIN product_details ON product_details.id = cart_items.product_details_id").
		Joins("JOIN products ON products.id = product_details.product_id").
		Joins("JOIN sizes ON sizes.id = product_details.size_id").
		Joins("JOIN brands ON brands.id = products.brand_id").
		Where("cart_items.cart_id = ?", cartID).
		Find(&viewCartDetails).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no cart items found for the given cart ID")
		}
		return nil, err
	}

	return viewCartDetails, nil
}
