package repositoryImpl

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/domain/entity"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db}
}

func (pr *ProductRepositoryImpl) CreateProduct(product *entity.Product) error {
	if err := pr.db.Create(product).Error; err != nil {
		return errors.Wrap(err, "failed to create product")
	}
	return nil
}

func (pr *ProductRepositoryImpl) ProductGetBySlug(slag string) (*entity.Product, error) {
	product := &entity.Product{}
	if err := pr.db.First(product, slag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, errors.Wrap(err, "failed to get product by slag")
	}
	return product, nil

}

func (pr *ProductRepositoryImpl) ProductGetByID(id uint) (*entity.Product, error) {
	product := &entity.Product{}
	if err := pr.db.First(product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, errors.Wrap(err, "failed to get product by ID")
	}
	return product, nil
}

func (pr *ProductRepositoryImpl) ProductGetByName(name string) (*entity.Product, error) {
	product := &entity.Product{}
	if err := pr.db.Where("name = ?", name).First(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, errors.Wrap(err, "failed to get product by name")
	}
	return product, nil
}

func (pr *ProductRepositoryImpl) ProductUpdate(product *entity.Product) error {
	if err := pr.db.Save(product).Error; err != nil {
		return errors.Wrap(err, "failed to update product")
	}
	return nil
}

func (pr *ProductRepositoryImpl) ProductDelete(product *entity.Product) error {
	toggle := !product.Deleted

	if err := pr.db.Model(product).Update("deleted", toggle).Error; err != nil {
		return errors.Wrap(err, "failed to soft delete product")
	}

	if err := pr.db.Model(product).Update("deleted_at", gorm.DeletedAt{}).Error; err != nil {
		return errors.Wrap(err, "failed to delete product")
	}
	return nil
}

func (pr *ProductRepositoryImpl) ProductListAll() ([]*entity.Product, error) {
	products := []*entity.Product{}
	if err := pr.db.Preload("Categories").Find(&products).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get all products")
	}
	return products, nil
}

func (pr *ProductRepositoryImpl) ProductListPaginated(offset, limit int) ([]*deliverymodels.ViewProductDetail, int, error) {
	products := []*deliverymodels.ViewProductDetail{}
	var totalCount int64

	// Perform the join with brands and sizes tables
	if err := pr.db.Table("products").
		Select("products.name, product_details.price, products.description, sizes.size AS size, brands.brand AS brand").
		Joins("JOIN product_details ON product_details.product_id = products.id").
		Joins("JOIN sizes ON sizes.id = product_details.size_id").
		Joins("JOIN brands ON brands.id = products.brand_id").
		Offset(offset).
		Limit(limit).
		Scan(&products).
		Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed to get paginated product list")
	}

	if err := pr.db.Model(&entity.Product{}).Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed to get total count of products")
	}

	return products, int(totalCount), nil
}

func (pr *ProductRepositoryImpl) CreateProductItem(productItem *entity.ProductDetails) error {
	if err := pr.db.Create(productItem).Error; err != nil {
		return errors.Wrap(err, "failed to create product item")
	}
	return nil
}

func (pr *ProductRepositoryImpl) ProductItemGetByID(id uint) (*entity.ProductDetails, error) {
	product := &entity.ProductDetails{}
	if err := pr.db.First(product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product item not found")
		}
		return nil, errors.Wrap(err, "failed to get product item by ID")
	}
	return product, nil
}

func (pr *ProductRepositoryImpl) ProductItemGetBySlug(slug string) (*entity.ProductDetails, error) {
	product := &entity.ProductDetails{}
	if err := pr.db.First(product, slug).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product item not found")
		}
		return nil, errors.Wrap(err, "failed to get product item by ID")
	}
	return product, nil
}

func (pr *ProductRepositoryImpl) ProductItemGetByProductIdAndSizeId(productId, sizeId uint) (*entity.ProductDetails, error) {
	product := &entity.ProductDetails{}
	if err := pr.db.Where("product_id = ? AND size_id = ?", productId, sizeId).First(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product item not found")
		}
		return nil, errors.Wrap(err, "failed to get product item by product and size ID")
	}
	return product, nil
}

func (pr *ProductRepositoryImpl) ProductItemUpdate(product *entity.ProductDetails) error {
	if product.ID == 0 {
		return errors.New("product ID cannot be zero")
	}

	existingProduct := &entity.ProductDetails{}
	if err := pr.db.First(existingProduct, product.ID).Error; err != nil {
		return errors.Wrap(err, "failed to find product item")
	}

	existingProduct.Price = product.Price
	existingProduct.Quantity = product.Quantity

	if err := pr.db.Save(existingProduct).Error; err != nil {
		return errors.Wrap(err, "failed to update product item")
	}

	return nil
}

func (pr *ProductRepositoryImpl) ProductItemDelete(product *entity.ProductDetails) error {
	if err := pr.db.Delete(product).Error; err != nil {
		return errors.Wrap(err, "failed to delete product item")
	}
	return nil
}

func (pr *ProductRepositoryImpl) ProductItemListAll() ([]*entity.ProductDetails, error) {
	products := []*entity.ProductDetails{}
	if err := pr.db.Find(&products).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get all product items")
	}
	return products, nil
}

func (pr *ProductRepositoryImpl) ProductItemListPaginated(offset, limit int) ([]*deliverymodels.ViewProductDetail, int, error) {
	products := []*deliverymodels.ViewProductDetail{}
	var totalCount int64

	if err := pr.db.Table("product_deatails").Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed to get paginated product item list")
	}

	if err := pr.db.Model(&entity.ProductDetails{}).Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed to get total count of product items")
	}

	return products, int(totalCount), nil
}

func (pr *ProductRepositoryImpl) GetProductsByBrandID(brandID uint, page int, limit int) ([]entity.Product, error) {
	var products []entity.Product
	offset := (page - 1) * limit
	result := pr.db.Table("products").
		Where("brand_id = ?", brandID).
		Offset(offset).
		Limit(limit).
		Find(&products)

	if result.Error != nil {
		fmt.Println("Failed to retrieve products:", result.Error)
		return nil, result.Error
	}

	return products, nil
}

func (pr *ProductRepositoryImpl) SearchProductByName(input string, page, limit int) ([]*entity.Product, error) {
	var products []*entity.Product
	offset := (page - 1) * limit
	result := pr.db.Where("name LIKE ?", input+"%").Offset(offset).Limit(limit).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (pr *ProductRepositoryImpl) GetProductDetailsByProductID(productID uint) ([]*entity.ProductDetails, error) {
	var productDetails []*entity.ProductDetails
	result := pr.db.Where("product_id = ?", productID).Find(&productDetails)
	if result.Error != nil {
		return nil, result.Error
	}
	return productDetails, nil
}
