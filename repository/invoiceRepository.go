package repository

import "70_Off/entity"

type InvoiceRepository interface {
	SaveInvoice(invoice *entity.Invoice) (*entity.Invoice, error)
	FindInvoiceById(id uint) (*entity.Invoice, error)
	UpdateInvoice(invoice *entity.Invoice) (*entity.Invoice, error)
	DeleteInvoice(invoice *entity.Invoice) error
}
