package services

// import (
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"

// 	deliverymodels "70_Off/delivery/deliveryModels"
// 	"70_Off/domain/entity"
// 	mockRepo "70_Off/domain/repository/mockRepo"

// 	"golang.org/x/crypto/bcrypt"
// )

// type MockRepositories struct {
// 	adminRepo     *mockRepo.MockAdminRepository
// 	userRepo      *mockRepo.MockUserRepository
// 	inventoryRepo *mockRepo.MockInventory
// 	orderRepo     *mockRepo.MockOrderRepository
// 	walletRepo    *mockRepo.MockWalletRepository
// 	couponRepo    *mockRepo.MockCouponRepository
// }

// func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
// 	return &MockRepositories{
// 		adminRepo:     mockRepo.NewMockAdminRepository(ctrl),
// 		userRepo:      mockRepo.NewMockUserRepository(ctrl),
// 		inventoryRepo: mockRepo.NewMockInventory(ctrl),
// 		orderRepo:     mockRepo.NewMockOrderRepository(ctrl),
// 		walletRepo:    mockRepo.NewMockWalletRepository(ctrl),
// 		couponRepo:    mockRepo.NewMockCouponRepository(ctrl),
// 	}
// }
// func TestAdminService_Create(t *testing.T) {
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

// 	admin := &deliverymodels.AdminSignupInputs{
// 		Name:        "John Doe",
// 		Email:       "john@example2.com",
// 		PhoneNumber: "1234567891",
// 		Password:    "password",
// 	}

// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
// 	expectedAdmin := &entity.Admin{
// 		Name:        admin.Name,
// 		Email:       admin.Email,
// 		PhoneNumber: admin.PhoneNumber,
// 		Password:    string(hashedPassword),
// 	}

// 	mockRepos.adminRepo.EXPECT().GetByEmail(admin.Email).Return(nil, nil)

// 	mockRepos.adminRepo.EXPECT().GetByPhoneNumber(admin.PhoneNumber).Return(nil, nil)

// 	mockRepos.adminRepo.EXPECT().Create(expectedAdmin).Return(nil)

// 	err := adminService.Create(admin)

// 	assert.Nil(t, err, "Expected no error")
// }
