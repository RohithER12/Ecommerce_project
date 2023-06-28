package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	deliverymodels "70_Off/delivery/deliveryModels"
	mockRepo "70_Off/domain/repository/mockRepo"
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
