package repositoryImpl

import (
	"70_Off/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (repo *UserRepositoryImpl) Create(user *entity.User) error {
	query := `INSERT INTO users (first_name, last_name, email, phone_number, username, password, block, verified) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result := repo.db.Exec(query, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Username, user.Password, user.Block, user.Verified)

	if result.Error != nil {
		return errors.New("failed to create user")
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows were affected")
	}

	return nil
}

// func (repo *UserRepositoryImpl) Create(users *entity.User) error {
// 	result := repo.db.Create(&users)
// 	if result.Error != nil {
// 		return errors.New("failed to create user")
// 	}
// 	return nil
// }

func (repo *UserRepositoryImpl) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetUserByUserName(username string) (*entity.User, error) {
	var user entity.User
	result := repo.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User
	result := repo.db.Where("phone_number = ?", phoneNumber).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) Update(user *entity.User) error {
	result := repo.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepositoryImpl) AddAddress(address *entity.Address) error {
	result := repo.db.Create(&address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepositoryImpl) GetAllAddressesByUserId(id uint) (*[]entity.Address, error) {
	var addresses []entity.Address
	result := repo.db.Preload("User").Where("user_id = ?", id).Find(&addresses)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(addresses) == 0 {
		return nil, errors.New("Address not added")
	}

	return &addresses, nil
}

func (repo *UserRepositoryImpl) GetAddressById(id uint) (*entity.Address, error) {

	var addresses entity.Address
	result := repo.db.Where("id = ?", id).First(&addresses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &addresses, nil
}
