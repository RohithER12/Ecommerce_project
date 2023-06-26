package repositoryImpl_test

import (
	"70_Off/domain/entity"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"70_Off/domain/repositoryImpl"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to set up test database: %s", err.Error())
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize gorm database: %s", err.Error())
	}

	return gormDB, mock
}

func teardownTestDB(db *gorm.DB, mock sqlmock.Sqlmock) {
	mock.ExpectClose()

}

func TestCreateAdmin(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	defer teardownTestDB(gormDB, mock)

	adminRepo := repositoryImpl.NewAdminRepositoryImpl(gormDB)

	admin := &entity.Admin{
		Name:        "John Doe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
		Password:    "password",
	}

	// Set up the expectations on the mock
	mock.ExpectExec("INSERT INTO admins").WithArgs(admin.Name, admin.Email, admin.PhoneNumber, admin.Password).WillReturnResult(sqlmock.NewResult(1, 1))

	err := adminRepo.Create(admin)
	assert.Nil(t, err, "Expected no error")

	// Assert that all expectations on the mock were met
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "Expected all expectations to be met")
}

func TestGetByID(t *testing.T) {
    gormDB, mock := setupTestDB(t)
    defer teardownTestDB(gormDB, mock)

    adminRepo := repositoryImpl.NewAdminRepositoryImpl(gormDB)

    expectedID := uint(1)
    expectedAdmin := &entity.Admin{
        Name:        "John Doe",
        Email:       "john@example.com",
        PhoneNumber: "1234567890",
        Password:    "password",
    }

    mock.ExpectQuery("^SELECT (.+) FROM admins WHERE id = \\$1$").
        WithArgs(expectedID).
        WillReturnRows(mock.NewRows([]string{"name", "email", "phonenumber", "password"}).
            AddRow(expectedAdmin.Name, expectedAdmin.Email, expectedAdmin.PhoneNumber, expectedAdmin.Password))

    admin, err := adminRepo.GetByID(expectedID)
    if err != nil {
        t.Errorf("Error retrieving admin: %s", err.Error())
    }
    fmt.Printf("Actual admin: %+v\n", admin)

    assert.Equal(t, expectedAdmin, admin)

    // Additional fix: Ensure that passing 0 as ID raises an error
    _, err = adminRepo.GetByID(0)
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


func TestGetByEmail(t *testing.T) {
	email := "john@example.com"
	expectedAdmin := &entity.Admin{
		Name:        "John Doe",
		Email:       email,
		PhoneNumber: "1234567890",
		Password:    "password",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to set up mock database: %s", err.Error())
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize GORM database: %s", err.Error())
	}

	adminRepo := repositoryImpl.NewAdminRepositoryImpl(gormDB)

	mock.ExpectQuery("^SELECT (.+) FROM admins WHERE email = \\$1$").
		WithArgs(email).
		WillReturnRows(mock.NewRows([]string{"name", "email", "phonenumber", "password"}).
			AddRow(expectedAdmin.Name, expectedAdmin.Email, expectedAdmin.PhoneNumber, expectedAdmin.Password))

	admin, err := adminRepo.GetByEmail(email)
	if err != nil {
		t.Errorf("Error retrieving admin: %s", err.Error())
	}

	assert.Equal(t, expectedAdmin, admin)

	_, err = adminRepo.GetByEmail("")
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

func TestGetByPhoneNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to set up mock database: %s", err.Error())
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize GORM database: %s", err.Error())
	}

	adminRepo := repositoryImpl.NewAdminRepositoryImpl(gormDB)

	expectedPhoneNumber := "1234567890"
	expectedAdmin := &entity.Admin{
		Name:        "John Doe",
		Email:       "john@example.com",
		PhoneNumber: expectedPhoneNumber,
		Password:    "password",
	}

	mock.ExpectQuery("^SELECT (.+) FROM admins WHERE phone_number = \\$1 LIMIT 1$").
		WithArgs(expectedPhoneNumber).
		WillReturnRows(mock.NewRows([]string{"name", "email", "phone_number", "password"}).
			AddRow(expectedAdmin.Name, expectedAdmin.Email, expectedAdmin.PhoneNumber, expectedAdmin.Password))

	admin, err := adminRepo.GetByPhoneNumber(expectedPhoneNumber)
	if err != nil {
		t.Errorf("Error retrieving admin: %s", err.Error())
	}
	fmt.Printf("Actual admin: %+v\n", admin)
	assert.Equal(t, expectedAdmin, admin)

	_, err = adminRepo.GetByPhoneNumber("")
	if err == nil {
		t.Error("Expected error when passing nil phoneNumber, but got no error")
	} else {
		fmt.Printf("Expected error: %s\n", err.Error())
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %s", err.Error())
	}

}

func TestUpdateAdmin(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	defer teardownTestDB(gormDB, mock)

	adminRepo := repositoryImpl.NewAdminRepositoryImpl(gormDB)

	admin := &entity.Admin{
		Model: gorm.Model{ID: 1},
		Name:  "John Doe",
		Email: "john@example.com",
		// ...
	}

	query := regexp.QuoteMeta("UPDATE admins SET name = $1, email = $2, phone_number = $3, password = $4 WHERE id = $5")

	mock.ExpectExec(query).
		WithArgs(admin.Name, admin.Email, admin.PhoneNumber, admin.Password, admin.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := adminRepo.Update(admin)
	assert.Nil(t, err, "Expected no error")

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "Expected all expectations to be met")
}
