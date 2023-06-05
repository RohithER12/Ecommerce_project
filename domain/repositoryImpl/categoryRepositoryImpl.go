package repositoryImpl

import (
	"70_Off/domain/entity"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db}
}

func (cr *CategoryRepositoryImpl) CreateSize(size *entity.Size) error {
	if err := cr.db.Create(size).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepositoryImpl) GetSizeByID(id uint) (*entity.Size, error) {
	size := &entity.Size{}
	if err := cr.db.First(size, id).Error; err != nil {
		return nil, err
	}
	return size, nil
}

func (cr *CategoryRepositoryImpl) GetSizeByName(name string) (*entity.Size, error) {
	size := &entity.Size{}
	if err := cr.db.Where("size = ?", name).First(size).Error; err != nil {
		return nil, err
	}
	return size, nil
}

func (cr *CategoryRepositoryImpl) UpdateSize(size *entity.Size) error {
	if err := cr.db.Save(size).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepositoryImpl) DeleteSize(size *entity.Size) error {
	if err := cr.db.Delete(size).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepositoryImpl) ListAllSizes() ([]*entity.Size, error) {
	sizes := []*entity.Size{}
	if err := cr.db.Find(&sizes).Error; err != nil {
		return nil, err
	}
	return sizes, nil
}

func (cr *CategoryRepositoryImpl) CreateBrand(brand *entity.Brand) error {
	if err := cr.db.Create(brand).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepositoryImpl) GetBrandByID(id uint) (*entity.Brand, error) {
	brand := &entity.Brand{}
	if err := cr.db.First(brand, id).Error; err != nil {
		return nil, err
	}
	return brand, nil
}

func (cr *CategoryRepositoryImpl) GetBrandByName(name string) (*entity.Brand, error) {
	brand := &entity.Brand{}
	if err := cr.db.Where("brand = ?", name).First(brand).Error; err != nil {
		return nil, err
	}
	return brand, nil
}

func (cr *CategoryRepositoryImpl) UpdateBrand(brand *entity.Brand) error {
	if err := cr.db.Save(brand).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepositoryImpl) DeleteBrand(brand *entity.Brand) error {
	if err := cr.db.Delete(brand).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepositoryImpl) ListAllBrands() ([]*entity.Brand, error) {
	brands := []*entity.Brand{}
	if err := cr.db.Find(&brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}
