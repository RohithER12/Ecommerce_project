package repositoryImpl_test

import (
	"70_Off/domain/entity"
	"70_Off/domain/repositoryImpl"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	defer teardownTestDB(gormDB, mock)

	userRepo := repositoryImpl.NewUserRepositoryImpl(gormDB)

	user := &entity.User{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "123456789",
		Username:    "johndoe",
		Password:    "password",
		Block:       false,
		Verified:    false,
	}

	// expected query and arguments
	mock.ExpectExec("INSERT INTO users").WithArgs(user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Username, user.Password, user.Block, user.Verified).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call actual func
	err := userRepo.Create(user)
	assert.Nil(t, err, "Expected no error")

	// Verify expectations
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "Expected all expectations to be met")

}

func TestGetByIDUser(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	defer teardownTestDB(gormDB, mock)

	userRepo := repositoryImpl.NewUserRepositoryImpl(gormDB)

	expectedID := uint(1)
	expectedUser := &entity.User{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "123456789",
		Username:    "johndoe",
		Password:    "password",
		Block:       false,
		Verified:    false,
	}

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE id = \\$1$").
		WithArgs(expectedID).
		WillReturnRows(mock.NewRows([]string{"first_name", "last_name", "email", "phone_number", "username", "password", "block", "verified"}).
			AddRow(expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, expectedUser.PhoneNumber, expectedUser.Username, expectedUser.Password, expectedUser.Block, expectedUser.Verified))

	user, err := userRepo.GetByID(expectedID)
	if err != nil {
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	fmt.Printf("Actual user: %+v\n", user)

	assert.Equal(t, expectedUser, user)

	_, err = userRepo.GetByID(0)
	if err == nil {
		t.Error("Expected error when passing nil ID, but got no error")
	} else {
		fmt.Printf("Expected error: %s\n", err.Error())
	}
	// Assert that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %s", err.Error())
	}
}

func TestGetByEmailUser(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	defer teardownTestDB(gormDB, mock)

	userRepo := repositoryImpl.NewUserRepositoryImpl(gormDB)

	expectedEmail := "john@example.com"
	expectedUser := &entity.User{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "123456789",
		Username:    "johndoe",
		Password:    "password",
		Block:       false,
		Verified:    false,
	}

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE email = \\$1$").
		WithArgs(expectedEmail).
		WillReturnRows(mock.NewRows([]string{"first_name", "last_name", "email", "phone_number", "username", "password", "block", "verified"}).
			AddRow(expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, expectedUser.PhoneNumber, expectedUser.Username, expectedUser.Password, expectedUser.Block, expectedUser.Verified))

	user, err := userRepo.GetByEmail(expectedEmail)
	if err != nil {
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	fmt.Printf("Actual user: %+v\n", user)

	assert.Equal(t, expectedUser, user)

	_, err = userRepo.GetByEmail("")
	if err == nil {
		t.Error("Expected error when passing nil email, but got no error")
	} else {
		fmt.Printf("Expected error: %s\n", err.Error())
	}
	// Assert that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %s", err.Error())
	}

}

func TestGetByPhoneNumberUser(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	defer teardownTestDB(gormDB, mock)

	userRepo := repositoryImpl.NewUserRepositoryImpl(gormDB)

	expectedPhoneNumber := "123456789"
	expectedUser := &entity.User{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "123456789",
		Username:    "johndoe",
		Password:    "password",
		Block:       false,
		Verified:    false,
	}

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE phone_number = \\$1$").
		WithArgs(expectedPhoneNumber).
		WillReturnRows(mock.NewRows([]string{"first_name", "last_name", "email", "phone_number", "username", "password", "block", "verified"}).
			AddRow(expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, expectedUser.PhoneNumber, expectedUser.Username, expectedUser.Password, expectedUser.Block, expectedUser.Verified))

	user, err := userRepo.GetByPhoneNumber(expectedPhoneNumber)
	if err != nil {
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	fmt.Printf("Actual user: %+v\n", user)

	assert.Equal(t, expectedUser, user)

	_, err = userRepo.GetByPhoneNumber("")
	if err == nil {
		t.Error("Expected error when passing nil email, but got no error")
	} else {
		fmt.Printf("Expected error: %s\n", err.Error())
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %s", err.Error())
	}
}

func TestGetByUserName(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	defer teardownTestDB(gormDB, mock)

	userRepo := repositoryImpl.NewUserRepositoryImpl(gormDB)

	expectedUsername := "johndoe"
	expectedUser := &entity.User{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "123456789",
		Username:    "johndoe",
		Password:    "password",
		Block:       false,
		Verified:    false,
	}

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE username = \\$1$").
		WithArgs(expectedUsername).
		WillReturnRows(mock.NewRows([]string{"first_name", "last_name", "email", "phone_number", "username", "password", "block", "verified"}).
			AddRow(expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, expectedUser.PhoneNumber, expectedUser.Username, expectedUser.Password, expectedUser.Block, expectedUser.Verified))

	user, err := userRepo.GetUserByUserName(expectedUsername)
	if err != nil {
		t.Errorf("Error retrieving user: %s", err.Error())
	}
	fmt.Printf("Actual user: %+v\n", user)

	assert.Equal(t, expectedUser, user)

	_, err = userRepo.GetUserByUserName("")
	if err == nil {
		t.Error("Expected error when passing nil username, but got no error")
	} else {
		fmt.Printf("Expected error: %s\n", err.Error())
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %s", err.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	defer teardownTestDB(gormDB, mock)

	userRepo := repositoryImpl.NewUserRepositoryImpl(gormDB)
	user := &entity.User{
		Model:       gorm.Model{ID: 1},
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "123456789",
		Username:    "johndoe",
		Password:    "password",
		Block:       false,
		Verified:    true,
	}

	query := regexp.QuoteMeta("UPDATE users SET first_name = $1, last_name = $2, email = $3, phone_number = $4, username = $5, password = $6, block = $7, verified = $8 WHERE id = $9")

	mock.ExpectExec(query).
		WithArgs(user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Username, user.Password, user.Block, user.Verified, user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := userRepo.Update(user)
	assert.Nil(t, err, "Expected no error")

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "Expected all expectations to be met")
}
