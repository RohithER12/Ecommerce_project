package repository

import "70_Off/entity"

type Inventory interface {
	CreateInventory(inventory *entity.Inventory) error
	GetInventoryByID(id uint) (*entity.Inventory, error)
	GetInventoryByProductItemID(id uint) (*entity.Inventory, error)
	UpdateInventory(inventory *entity.Inventory) error
	DeleteInventory(inventory *entity.Inventory) error
	ReduceQuantity(inventory *entity.Inventory, quantity uint) error
}
