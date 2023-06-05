package repositoryImpl

import (
	"70_Off/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type OtpValidationRepo struct {
	db *gorm.DB
}

func NewOtpValidationRepoImpl(db *gorm.DB) *OtpValidationRepo {
	return &OtpValidationRepo{db: db}
}

func (r *OtpValidationRepo) Create(otpValidation *entity.OtpValidation) (*entity.OtpValidation, error) {
	if err := r.db.Create(&otpValidation).Error; err != nil {
		return nil, errors.New("failed to create OtpValidation")
	}
	return otpValidation, nil
}

func (r *OtpValidationRepo) FindByKey(key string) (*entity.OtpValidation, error) {
	var otpValidation entity.OtpValidation
	if err := r.db.Where("key = ?", key).First(&otpValidation).Error; err != nil {
		return nil, errors.New("failed to find OtpValidation by key")
	}
	return &otpValidation, nil
}
