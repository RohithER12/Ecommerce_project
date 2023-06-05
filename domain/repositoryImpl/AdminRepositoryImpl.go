package repositoryImpl

import (
	"70_Off/domain/entity"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepositoryImpl(db *gorm.DB) *AdminRepositoryImpl {
	return &AdminRepositoryImpl{db}
}

func (ar *AdminRepositoryImpl) Create(admin *entity.Admin) error {
	return ar.db.Create(admin).Error
}

func (ar *AdminRepositoryImpl) GetByID(id uint) (*entity.Admin, error) {
	admin := &entity.Admin{}
	if err := ar.db.First(admin, id).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (ar *AdminRepositoryImpl) GetByEmail(email string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	if err := ar.db.Where("email = ?", email).First(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (ar *AdminRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	if err := ar.db.Where("phone_number = ?", phoneNumber).First(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (ar *AdminRepositoryImpl) Update(admin *entity.Admin) error {
	return ar.db.Save(admin).Error
}

func (ar *AdminRepositoryImpl) Block(user *entity.User) error {
	user.Block = true
	if err := ar.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AdminRepositoryImpl) Unblock(user *entity.User) error {
	user.Block = false
	if err := ar.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
