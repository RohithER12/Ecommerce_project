package repository

import (
	"70_Off/entity"
)

type Category interface {
	CreateSize(size *entity.Size) error
	GetSizeByID(id uint) (*entity.Size, error)
	GetSizeByName(name string) (*entity.Size, error)
	UpdateSize(size *entity.Size) error
	DeleteSize(size *entity.Size) error
	ListAllSizes() ([]*entity.Size, error)
	CreateBrand(brand *entity.Brand) error
	GetBrandByID(id uint) (*entity.Brand, error)
	GetBrandByName(name string) (*entity.Brand, error)
	UpdateBrand(brand *entity.Brand) error
	DeleteBrand(brand *entity.Brand) error
	ListAllBrands() ([]*entity.Brand, error)
}
