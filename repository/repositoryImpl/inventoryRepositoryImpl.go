package repositoryImpl

import (
	"70_Off/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type InventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepositoryImpl(db *gorm.DB) *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{db}
}

func (ir *InventoryRepositoryImpl) CreateInventory(inventory *entity.Inventory) error {
	if err := ir.db.Create(inventory).Error; err != nil {
		return errors.Wrap(err, "failed to create inventory")
	}
	return nil
}

func (ir *InventoryRepositoryImpl) GetInventoryByID(id uint) (*entity.Inventory, error) {
	inventory := &entity.Inventory{}
	if err := ir.db.First(inventory, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found")
		}
		return nil, errors.Wrap(err, "failed to get inventory by ID")
	}
	return inventory, nil
}

func (ir *InventoryRepositoryImpl) GetInventoryByProductItemID(id uint) (*entity.Inventory, error) {
	inventory := &entity.Inventory{}
	if err := ir.db.Where("product_deatails_id = ?", id).First(inventory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found")
		}
		return nil, errors.Wrap(err, "failed to get inventory by product item ID")
	}
	return inventory, nil
}

func (ir *InventoryRepositoryImpl) UpdateInventory(inventory *entity.Inventory) error {
	if err := ir.db.Save(inventory).Error; err != nil {
		return errors.Wrap(err, "failed to update inventory")
	}
	return nil
}

func (ir *InventoryRepositoryImpl) DeleteInventory(inventory *entity.Inventory) error {
	if err := ir.db.Delete(inventory).Error; err != nil {
		return errors.Wrap(err, "failed to delete inventory")
	}
	return nil
}

func (ir *InventoryRepositoryImpl) ReduceQuantity(inventory *entity.Inventory, quantity uint) error {
	if inventory.Quantity < quantity {
		return errors.New("insufficient quantity")
	}
	inventory.Quantity -= quantity
	if err := ir.db.Save(inventory).Error; err != nil {
		return errors.Wrap(err, "failed to reduce quantity")
	}
	return nil
}
