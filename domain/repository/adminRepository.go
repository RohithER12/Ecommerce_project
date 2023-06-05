package repository

import "70_Off/domain/entity"

type AdminRepository interface {
	Create(admin *entity.Admin) error
	GetByID(id uint) (*entity.Admin, error)
	GetByEmail(email string) (*entity.Admin, error)
	GetByPhoneNumber(phoneNumber string) (*entity.Admin, error)
	Update(admin *entity.Admin) error
	// Block(user *entity.User) error
	// Unblock(user *entity.User) error
}
