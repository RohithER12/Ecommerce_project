package repositoryImpl

import (
	"70_Off/entity"

	"gorm.io/gorm"
)

type InvoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepo(db *gorm.DB) *InvoiceRepo {
	return &InvoiceRepo{db: db}
}

func (r *InvoiceRepo) SaveInvoice(invoice *entity.Invoice) (*entity.Invoice, error) {
	if err := r.db.Create(&invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *InvoiceRepo) FindInvoiceById(id uint) (*entity.Invoice, error) {
	var invoice entity.Invoice
	if err := r.db.Preload("Address").Preload("Product").Preload("Payment").First(&invoice, id).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *InvoiceRepo) UpdateInvoice(invoice *entity.Invoice) (*entity.Invoice, error) {
	if err := r.db.Save(&invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *InvoiceRepo) DeleteInvoice(invoice *entity.Invoice) error {
	if err := r.db.Delete(&invoice).Error; err != nil {
		return err
	}
	return nil
}
