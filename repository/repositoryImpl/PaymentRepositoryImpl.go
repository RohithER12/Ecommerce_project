package repositoryImpl

import (
	"70_Off/entity"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) SavePayment(payment *entity.Payment) (*entity.Payment, error) {
	if err := r.db.Create(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepo) FindPaymentById(id uint) (*entity.Payment, error) {
	var payment entity.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepo) UpdatePayment(payment *entity.Payment) (*entity.Payment, error) {
	if err := r.db.Save(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepo) DeletePayment(payment *entity.Payment) error {
	if err := r.db.Delete(&payment).Error; err != nil {
		return err
	}
	return nil
}
