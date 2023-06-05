package services

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/delivery/helpers"
	"70_Off/domain/entity"
	"70_Off/domain/repository"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Product interface {
	ProductCreate(adminId uint, inputs *deliverymodels.AddProductInputs) error
	ProductGetByID(id uint) (*entity.Product, error)
	ProductUpdate(adminId uint, input *deliverymodels.UpdateProductInput) error
	ProductDelete(product *entity.Product) error
	ProductListAll() ([]*entity.Product, error)
	ProductListPaginated(offset, limit int) ([]*deliverymodels.ViewProductDetail, int, error)
	ProductCreateSize(size *entity.Size) error
	ProductUpdateSize(ProductID uint, small, large uint, medium uint) (*entity.Size, error)
	ProductGetByName(name string) (*entity.Product, error)
	CreateProductDeatails(adminID uint, input *deliverymodels.AddProductsDeatils) error
	ProductDetailsUpdate(adminId uint, input *deliverymodels.UpdateProductsDeatils) error
}

type ProductService struct {
	productRepo   repository.Product
	categoryRepo  repository.Category
	inventoryRepo repository.Inventory
	orderRepo     repository.OrderRepository
}

func NewProductService(productRepo repository.Product, categoryRepo repository.Category, inventoryRepo repository.Inventory, orderRepo repository.OrderRepository) *ProductService {
	return &ProductService{productRepo: productRepo, categoryRepo: categoryRepo, inventoryRepo: inventoryRepo, orderRepo: orderRepo}

}

func (ps *ProductService) ProductGetByID(id uint) (*entity.Product, error) {
	product, err := ps.productRepo.ProductGetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) ProductGetByName(name string) (*entity.Product, error) {
	product, err := ps.productRepo.ProductGetByName(name)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) ProductUpdate(adminId uint, input *deliverymodels.UpdateProductInput) error {
	getProduct, err := ps.productRepo.ProductGetByID(input.ProductId)
	if err != nil {
		return errors.New("There is no product in the database")
	}

	product := &entity.Product{
		Model: gorm.Model{
			ID: getProduct.ID,
		},
		Name:        input.Name,
		Description: input.Description,
		AdminID:     adminId,
		BrandID:     input.BrandId,
	}

	if err := ps.productRepo.ProductUpdate(product); err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) ProductDetailsUpdate(adminId uint, input *deliverymodels.UpdateProductsDeatils) error {

	getProductItem, err := ps.productRepo.ProductItemGetByProductIdAndSizeId(input.ProductId, input.SizeId)
	if err != nil {
		return errors.New("There is no productitem in database")

	}

	var productItem = entity.ProductDetails{
		Model: gorm.Model{
			ID: getProductItem.ID,
		},
		Price:     input.Price,
		ProductID: input.ProductId,
		SizeID:    input.SizeId,
		Quantity:  getProductItem.Quantity + input.Quantity,
	}

	if err := ps.productRepo.ProductItemUpdate(&productItem); err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) ProductDeleteToggle(id uint) error {
	product, err := ps.productRepo.ProductGetByID(id)
	if err != nil {
		return err
	}
	deleteToggle := !product.Deleted
	product.Deleted = deleteToggle
	if err := ps.productRepo.ProductUpdate(product); err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) ProductDelete(id uint) error {
	product, err := ps.productRepo.ProductGetByID(id)
	if err != nil {
		return err
	}
	if err := ps.productRepo.ProductDelete(product); err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) ProductListAll() ([]*entity.Product, error) {
	products, err := ps.productRepo.ProductListAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) ProductListPaginated(offset, limit int) ([]*deliverymodels.ViewProductDetail, int, error) {

	products, totalCount, err := ps.productRepo.ProductListPaginated(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

func (ps *ProductService) ProductCreate(adminId uint, input *deliverymodels.AddProductInputs) error {

	name, err := helpers.GenerateUniqueID(30)
	if err != nil {
		return err
	}
	slug := slug.Make(name)
	var product = entity.Product{
		Name:        input.Name,
		Description: input.Description,
		AdminID:     adminId,
		BrandID:     input.BrandId,
		Slug:        slug,
	}
	if err := ps.productRepo.CreateProduct(&product); err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) CreateProductDeatails(adminID uint, input *deliverymodels.AddProductsDeatils) error {

	productDetail, err := ps.productRepo.ProductItemGetByProductIdAndSizeId(input.ProductId, input.SizeId)
	if err != nil {
		fmt.Println("ProductDetail not found", err)
	}
	if productDetail != nil {
		return errors.New("Already Exist, Try Update")
	}
	forCreate := entity.ProductDetails{
		AdminID:   adminID,
		Price:     input.Price,
		SizeID:    input.SizeId,
		ProductID: input.ProductId,
		Quantity:  input.Quantity,
	}

	if err := ps.productRepo.CreateProductItem(&forCreate); err != nil {
		return err
	}

	return nil
}
