package repository

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/domain/entity"
)

type Product interface {
	CreateProduct(product *entity.Product) error
	ProductGetByID(id uint) (*entity.Product, error)
	ProductGetByName(name string) (*entity.Product, error)
	ProductUpdate(product *entity.Product) error
	ProductDelete(product *entity.Product) error
	ProductListAll() ([]*entity.Product, error)
	ProductListPaginated(offset, limit int) ([]*deliverymodels.ViewProductDetail, int, error)
	GetProductsByBrandID(brandID uint, page int, limit int) ([]entity.Product, error)
	SearchProductByName(input string, page, limit int) ([]*entity.Product, error)
	ProductGetBySlug(slug string) (*entity.Product, error)

	CreateProductItem(product *entity.ProductDetails) error
	ProductItemGetByID(id uint) (*entity.ProductDetails, error)
	ProductItemGetByProductIdAndSizeId(productId, sizeId uint) (*entity.ProductDetails, error)
	ProductItemUpdate(product *entity.ProductDetails) error
	ProductItemDelete(product *entity.ProductDetails) error
	GetProductDetailsByProductID(productID uint) ([]*entity.ProductDetails, error)
	// ProdutItemGetBySlug(slug string) (*entity.ProductDetails, error)
}
