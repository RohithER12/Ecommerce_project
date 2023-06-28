package repositoryImpl

import (
	"70_Off/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepositoryImpl(db *gorm.DB) *AdminRepositoryImpl {
	return &AdminRepositoryImpl{db}
}

func (ar *AdminRepositoryImpl) Create(admin *entity.Admin) error {
	query := `
		INSERT INTO admins (name, email, phonenumber, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`

	err := ar.db.Exec(query, admin.Name, admin.Email, admin.PhoneNumber, admin.Password).Error
	return err
}

// func (ar *AdminRepositoryImpl) Create(admin *entity.Admin) error {
// 	result := ar.db.Create(&admin)
// 	if result.Error != nil {
// 		return errors.New("failed to create admin")
// 	}
// 	return nil
// }

func (ar *AdminRepositoryImpl) GetByID(id uint) (*entity.Admin, error) {
	admin := &entity.Admin{}
	query := `
		SELECT name, email, phonenumber, password FROM admins WHERE id = $1
	`
	rows, err := ar.db.Raw(query, id).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name, email, phoneNumber, password string
		if err := rows.Scan(&name, &email, &phoneNumber, &password); err != nil {
			return nil, err
		}

		admin.Name = name
		admin.Email = email
		admin.PhoneNumber = phoneNumber
		admin.Password = password
	}

	return admin, nil
}

// func (ar *AdminRepositoryImpl) GetByID(id uint) (*entity.Admin, error) {
// 	admin := &entity.Admin{}
// 	if err := ar.db.First(admin, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return admin, nil
// }

func (ar *AdminRepositoryImpl) GetByEmail(email string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	query := `
		SELECT name, email, phonenumber, password FROM admins WHERE email = $1
	`
	rows, err := ar.db.Raw(query, email).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&admin.Name, &admin.Email, &admin.PhoneNumber, &admin.Password); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return admin, nil
}

// func (ar *AdminRepositoryImpl) GetByEmail(email string) (*entity.Admin, error) {
// 	admin := &entity.Admin{}
// 	if err := ar.db.Where("email = ?", email).First(admin).Error; err != nil {
// 		return nil, err
// 	}
// 	return admin, nil
// }

func (ar *AdminRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	query := `SELECT * FROM admins WHERE phone_number = $1 LIMIT 1`
	if err := ar.db.Raw(query, phoneNumber).Scan(admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return admin, nil
}

// func (ar *AdminRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*entity.Admin, error) {
// 	admin := &entity.Admin{}
// 	if err := ar.db.Where("phone_number = ?", phoneNumber).First(admin).Error; err != nil {
// 		return nil, err
// 	}
// 	return admin, nil
// }

func (ar *AdminRepositoryImpl) Update(admin *entity.Admin) error {
	query := "UPDATE admins SET name = ?, email = ?, phone_number = ?, password = ? WHERE id = ?"

	result := ar.db.Exec(query, admin.Name, admin.Email, admin.PhoneNumber, admin.Password, admin.ID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows were affected")
	}

	return nil
}

// func (ar *AdminRepositoryImpl) Update(admin *entity.Admin) error {
// 	return ar.db.Save(admin).Error
// }

func (ar *AdminRepositoryImpl) Block(user *entity.User) error {
	user.Block = true
	if err := ar.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AdminRepositoryImpl) Unblock(user *entity.User) error {
	user.Block = false
	if err := ar.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
