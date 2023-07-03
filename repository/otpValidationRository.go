package repository

import "70_Off/entity"

type OtpValidationRepository interface {
	Create(OtpValidation *entity.OtpValidation) (*entity.OtpValidation, error)
	FindByKey(key string) (*entity.OtpValidation, error)
}
