package repositoryImpl

import (
	"70_Off/entity"
	"errors"
	"fmt"

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

//	func (repo *UserRepositoryImpl) Create(users *entity.User) error {
//		result := repo.db.Create(&users)
//		if result.Error != nil {
//			return errors.New("failed to create user")
//		}
//		return nil
//	}
func (repo *UserRepositoryImpl) GetByID(id uint) (*entity.User, error) {
	user := &entity.User{}
	query := `
		SELECT first_name, last_name, email, phone_number, username, password, block, verified FROM users WHERE id = $1
	`
	rows, err := repo.db.Raw(query, id).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var firstName, lastName, email, phoneNumber, username, password string
		var block, verified bool
		if err := rows.Scan(&firstName, &lastName, &email, &phoneNumber, &username, &password, &block, &verified); err != nil {
			return nil, err
		}

		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.PhoneNumber = phoneNumber
		user.Username = username
		user.Password = password
		user.Block = block
		user.Verified = verified
	}

	return user, nil
}

// func (repo *UserRepositoryImpl) GetByID(id uint) (*entity.User, error) {
// 	var user entity.User
// 	result := repo.db.First(&user, id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &user, nil
// }

func (repo *UserRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := `
	SELECT first_name, last_name, email, phone_number, username, password, block, verified FROM users WHERE email = $1
`
	rows, err := repo.db.Raw(query, email).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var firstName, lastName, email, phoneNumber, username, password string
		var block, verified bool
		if err := rows.Scan(&firstName, &lastName, &email, &phoneNumber, &username, &password, &block, &verified); err != nil {
			return nil, err
		}

		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.PhoneNumber = phoneNumber
		user.Username = username
		user.Password = password
		user.Block = block
		user.Verified = verified
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) GetUserByUserName(username string) (*entity.User, error) {
	var user entity.User
	query := `
	SELECT first_name, last_name, email, phone_number, username, password, block, verified FROM users WHERE username = $1
`
	rows, err := repo.db.Raw(query, username).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var firstName, lastName, email, phoneNumber, username, password string
		var block, verified bool
		if err := rows.Scan(&firstName, &lastName, &email, &phoneNumber, &username, &password, &block, &verified); err != nil {
			return nil, err
		}

		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.PhoneNumber = phoneNumber
		user.Username = username
		user.Password = password
		user.Block = block
		user.Verified = verified
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User
	query := `
	SELECT first_name, last_name, email, phone_number, username, password, block, verified FROM users WHERE phone_number = $1`
	rows, err := repo.db.Raw(query, phoneNumber).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var firstName, lastName, email, phoneNumber, username, password string
		var block, verified bool
		if err := rows.Scan(&firstName, &lastName, &email, &phoneNumber, &username, &password, &block, &verified); err != nil {
			return nil, err
		}

		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.PhoneNumber = phoneNumber
		user.Username = username
		user.Password = password
		user.Block = block
		user.Verified = verified
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) Update(user *entity.User) error {
	query := "UPDATE users SET first_name = ?, last_name = ?, email = ?, phone_number = ?, username = ?, password = ?, block = ?, verified = ? WHERE id = ?"

	result := repo.db.Exec(query, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Username, user.Password, user.Block, user.Verified, user.ID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows were affected")
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
		return nil, errors.New("address not added")
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

// for testing
