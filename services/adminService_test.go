package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/entity"
	mockRepo "70_Off/repository/mockRepo"
)

type MockRepositories struct {
	adminRepo     *mockRepo.MockAdminRepository
	userRepo      *mockRepo.MockUserRepository
	inventoryRepo *mockRepo.MockInventory
	orderRepo     *mockRepo.MockOrderRepository
	walletRepo    *mockRepo.MockWalletRepository
	couponRepo    *mockRepo.MockCouponRepository
}

func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
	return &MockRepositories{
		adminRepo:     mockRepo.NewMockAdminRepository(ctrl),
		userRepo:      mockRepo.NewMockUserRepository(ctrl),
		inventoryRepo: mockRepo.NewMockInventory(ctrl),
		orderRepo:     mockRepo.NewMockOrderRepository(ctrl),
		walletRepo:    mockRepo.NewMockWalletRepository(ctrl),
		couponRepo:    mockRepo.NewMockCouponRepository(ctrl),
	}
}
func TestAdminService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepos := NewMockRepositories(ctrl)

	adminService := NewAdminService(
		mockRepos.adminRepo,
		mockRepos.userRepo,
		mockRepos.inventoryRepo,
		mockRepos.orderRepo,
		mockRepos.walletRepo,
		mockRepos.couponRepo,
	)

	admin := &deliverymodels.AdminSignupInputs{
		Name:        "John Doe",
		Email:       "john@example2.com",
		PhoneNumber: "1234567891",
		Password:    "password",
	}

	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	// expectedAdmin := &entity.Admin{
	// 	Name:        admin.Name,
	// 	Email:       admin.Email,
	// 	PhoneNumber: admin.PhoneNumber,
	// 	Password:    gomock.Any().String(),
	// }

	mockRepos.adminRepo.EXPECT().GetByEmail(admin.Email).Return(nil, errors.New("invalid email"))

	mockRepos.adminRepo.EXPECT().GetByPhoneNumber(admin.PhoneNumber).Return(nil, errors.New("invalid phone number"))

	mockRepos.adminRepo.EXPECT().Create(gomock.Any()).Return(nil)

	err := adminService.Create(admin)

	assert.Nil(t, err, nil)
}

func TestAdminService_CreateEmailAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepos := NewMockRepositories(ctrl)

	adminService := NewAdminService(
		mockRepos.adminRepo,
		mockRepos.userRepo,
		mockRepos.inventoryRepo,
		mockRepos.orderRepo,
		mockRepos.walletRepo,
		mockRepos.couponRepo,
	)

	admin := &deliverymodels.AdminSignupInputs{
		Name:        "John Doe",
		Email:       "john@example2.com",
		PhoneNumber: "1234567891",
		Password:    "password",
	}
	mockRepos.adminRepo.EXPECT().GetByEmail(admin.Email).Return(nil, nil)
	mockRepos.adminRepo.EXPECT().Create(gomock.Any()).Return(errors.New("email already exists")).AnyTimes()

	err := adminService.Create(admin)
	assert.Error(t, err, "email already exists")
}

func TestAdminService_CreatePhoneNumberAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepos := NewMockRepositories(ctrl)

	adminService := NewAdminService(
		mockRepos.adminRepo,
		mockRepos.userRepo,
		mockRepos.inventoryRepo,
		mockRepos.orderRepo,
		mockRepos.walletRepo,
		mockRepos.couponRepo,
	)

	admin := &deliverymodels.AdminSignupInputs{
		Name:        "John Doe",
		Email:       "john@example2.com",
		PhoneNumber: "1234567891",
		Password:    "password",
	}
	mockRepos.adminRepo.EXPECT().GetByEmail(admin.Email).Return(nil, errors.New("invalid email"))
	mockRepos.adminRepo.EXPECT().GetByPhoneNumber(admin.PhoneNumber).Return(nil, nil)
	mockRepos.adminRepo.EXPECT().Create(gomock.Any()).Return(errors.New("phoneNumber already exists")).AnyTimes()

	err := adminService.Create(admin)
	assert.Error(t, err, "phoneNumber already exists")
}

func TestAdminService_AdminValidateLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepos := NewMockRepositories(ctrl)

	adminService := NewAdminService(
		mockRepos.adminRepo,
		mockRepos.userRepo,
		mockRepos.inventoryRepo,
		mockRepos.orderRepo,
		mockRepos.walletRepo,
		mockRepos.couponRepo,
	)

	// Valid credentials test case
	adminLoginInput := &deliverymodels.AdmionLoginInputs{
		Email:    "john@example2.com",
		Password: "password",
	}

	dummyAdmin := &entity.Admin{
		Email:    adminLoginInput.Email,
		Password: "hashedPassword",
	}
	mockRepos.adminRepo.EXPECT().GetByEmail(adminLoginInput.Email).Return(dummyAdmin, nil)

	admin, err := adminService.AdminValidateLogin(adminLoginInput)

	assert.Nil(t, admin)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid credentials")

	// Invalid password test case
	invalidPasswordInput := &deliverymodels.AdmionLoginInputs{
		Email:    "john@example2.com",
		Password: "invalidPassword",
	}

	dummyAdmin.Password = "hashedPassword"
	mockRepos.adminRepo.EXPECT().GetByEmail(invalidPasswordInput.Email).Return(dummyAdmin, nil)

	admin, err = adminService.AdminValidateLogin(invalidPasswordInput)

	assert.Nil(t, admin)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid credentials")
}

func TestAdminService_UserBlockToggle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepos := NewMockRepositories(ctrl)
	mockUserRepo := mockRepo.NewMockUserRepository(ctrl)

	adminService := NewAdminService(
		mockRepos.adminRepo,
		mockUserRepo,
		mockRepos.inventoryRepo,
		mockRepos.orderRepo,
		mockRepos.walletRepo,
		mockRepos.couponRepo,
	)

	userID := uint(1)
	dummyUser := &entity.User{
		Model: gorm.Model{
			ID: userID,
		},
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
		Username:    "johndoe",
		Password:    "password",
		Block:       false,
		Verified:    false,
	}

	mockUserRepo.EXPECT().GetByID(userID).Return(dummyUser, nil)

	mockUserRepo.EXPECT().Update(dummyUser).Return(nil)

	err := adminService.UserBlockToggle(userID)

	assert.Nil(t, err)
	assert.True(t, dummyUser.Block)
}

// func TestAdminService_OrderReport(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepos := NewMockRepositories(ctrl)

// 	adminService := NewAdminService(
// 		mockRepos.adminRepo,
// 		mockRepos.userRepo,
// 		mockRepos.inventoryRepo,
// 		mockRepos.orderRepo,
// 		mockRepos.walletRepo,
// 		mockRepos.couponRepo,
// 	)

// 	n := uint(7) // No of days

// 	// Mock GetOrderStatistics - given days
// 	startDate := time.Now().AddDate(0, 0, -int(n-1)).Round(0)
// 	endDate := time.Now().Round(0)
// 	mockRepos.orderRepo.EXPECT().GetOrderStatistics(gomock.Eq(startDate), gomock.Eq(endDate)).Return(int64(10), 500.0, nil)

// 	// Mock the GetOrderStatistics  - last 7 days
// 	last7DaysStartDate := time.Now().AddDate(0, 0, -6).Round(0)
// 	mockRepos.orderRepo.EXPECT().GetOrderStatistics(gomock.Eq(last7DaysStartDate), gomock.Any()).Return(int64(5), 250.0, nil)

// 	// Mock the GetOrderStatistics - last month
// 	lastMonthStartDate := time.Now().AddDate(0, -1, 0).Round(0)
// 	mockRepos.orderRepo.EXPECT().GetOrderStatistics(gomock.Eq(lastMonthStartDate), gomock.Any()).Return(int64(20), 1000.0, nil)

// 	// Mock the GetOrderStatistics - last year
// 	lastYearStartDate := time.Now().AddDate(-1, 0, 0).Round(0)
// 	mockRepos.orderRepo.EXPECT().GetOrderStatistics(gomock.Eq(lastYearStartDate), gomock.Any()).Return(int64(100), 5000.0, nil)

// 	report, err := adminService.OrderReport(n)

// 	// Assert the results
// 	assert.Nil(t, err)
// 	assert.Equal(t, uint(10), report.NoOfOrders)
// 	assert.Equal(t, 500.0, report.TotalAmount)
// 	assert.Equal(t, uint(5), report.WeeklyOrders)
// 	assert.Equal(t, 250.0, report.WeeklyTotalAmount)
// 	assert.Equal(t, uint(20), report.MonthlyOrders)
// 	assert.Equal(t, 1000.0, report.MonthlyTotalAmount)
// 	assert.Equal(t, uint(100), report.YearlyOrders)
// 	assert.Equal(t, 5000.0, report.YearlyTotalAmount)
// }
