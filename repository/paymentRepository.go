package repository

import "70_Off/entity"

type PaymentRepository interface {
	SavePayment(payment *entity.Payment) (*entity.Payment, error)
	FindPaymentById(id uint) (*entity.Payment, error)
	UpdatePayment(payment *entity.Payment) (*entity.Payment, error)
	DeletePayment(payment *entity.Payment) error
}
