package repositoryImpl_test

import (
	"70_Off/domain/entity"
	"70_Off/domain/repositoryImpl"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer mockDB.Close()

	// Create a new GORM DB instance with the mock DB connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database connection: %s", err)
	}

	// Create a new UserRepositoryImpl instance with the mock DB
	userRepo := repositoryImpl.NewUserRepositoryImpl(db)

	// Set up the expected query and arguments
	expectedQuery := `INSERT INTO users (first_name, last_name, email, phone_number, username, password, block, verified) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Set up the mock expectations
	mock.ExpectExec(expectedQuery).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Create a user entity to pass to the Create method
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

	// Call the Create method
	err = userRepo.Create(user)

	// Check for any error
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}
