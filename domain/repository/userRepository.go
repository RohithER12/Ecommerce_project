package repository

import "70_Off/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetUserByUserName(username string) (*entity.User, error)
	GetByPhoneNumber(phoneNumber string) (*entity.User, error)
	Update(user *entity.User) error
	AddAddress(address *entity.Address) error
	GetAllAddressesByUserId(id uint) (*[]entity.Address, error)
	GetAddressById(id uint) (*entity.Address, error)
}
