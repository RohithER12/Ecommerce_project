package services

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/domain/entity"
	"70_Off/domain/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	Create(admin *deliverymodels.AdminSignupInputs) error
	GetByID(id uint) (*entity.Admin, error)
	GetByEmail(email string) (*entity.Admin, error)
	GetByPhoneNumber(phoneNumber string) (*entity.Admin, error)
	Update(admin *entity.Admin) (*entity.Admin, error)
	AdminValidateLogin(adminLoginInput *deliverymodels.AdmionLoginInputs) (*entity.Admin, error)
	UserBlockToggle(id uint) error
	OrderReport(days uint) (deliverymodels.OrderReport, error)
	UpdatingOrderStatus(orderId uint) (*entity.Orders, error)
	UpdatingReturnOrderStatus(orderId uint) error
	AddCoupon(coupon entity.Coupon) error
}

type adminService struct {
	adminRepo     repository.AdminRepository
	userRepo      repository.UserRepository
	inventoryRepo repository.Inventory
	orderRepo     repository.OrderRepository
	walletRepo    repository.WalletRepository
	couponRepo    repository.CouponRepository
}

func NewAdminService(adminRepo repository.AdminRepository, userRepo repository.UserRepository,
	inventoryRepo repository.Inventory, orderRepo repository.OrderRepository,
	walletRepo repository.WalletRepository, couponRepo repository.CouponRepository) *adminService {

	return &adminService{adminRepo: adminRepo, userRepo: userRepo,
		inventoryRepo: inventoryRepo, orderRepo: orderRepo, walletRepo: walletRepo, couponRepo: couponRepo}
}

func (s *adminService) Create(admin *deliverymodels.AdminSignupInputs) error {

	email := admin.Email
	phoneNumber := admin.PhoneNumber

	if _, err := s.adminRepo.GetByEmail(email); err == nil {
		return errors.New("Email already exists")
	}

	if _, err := s.adminRepo.GetByPhoneNumber(phoneNumber); err == nil {
		return errors.New("PhoneNumber already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	toCrateAdmin := entity.Admin{
		Name:        admin.Name,
		Email:       admin.Email,
		PhoneNumber: admin.PhoneNumber,
		Password:    string(hashedPassword),
	}

	err1 := s.adminRepo.Create(&toCrateAdmin)
	if err1 != nil {
		return err1
	}
	return nil
}

func (s *adminService) AdminValidateLogin(adminLoginInput *deliverymodels.AdmionLoginInputs) (*entity.Admin, error) {

	email := adminLoginInput.Email
	password := adminLoginInput.Password

	admin, err := s.adminRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid credentials")
	}

	return admin, nil
}

func (s *adminService) GetByID(id uint) (*entity.Admin, error) {

	admin, err := s.adminRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *adminService) GetByEmail(email string) (*entity.Admin, error) {

	admin, err := s.adminRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *adminService) GetByPhoneNumber(phoneNumber string) (*entity.Admin, error) {

	admin, err := s.adminRepo.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *adminService) Update(admin *entity.Admin) (*entity.Admin, error) {

	err := s.adminRepo.Update(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *adminService) UserBlockToggle(id uint) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	blockToggle := !user.Block
	user.Block = blockToggle
	err1 := s.userRepo.Update(user)
	if err1 != nil {
		return err1
	}
	return nil
}

func (s *adminService) OrderReport(n uint) (deliverymodels.OrderReport, error) {
	var report deliverymodels.OrderReport
	currentTime := time.Now()
	startDate := currentTime.AddDate(0, 0, -int(n-1))
	endDate := currentTime
	last7DaysStartDate := currentTime.AddDate(0, 0, -6)
	lastMonthStartDate := currentTime.AddDate(0, -1, 0)
	lastYearStartDate := currentTime.AddDate(-1, 0, 0)

	orderCount, totalAmount, err := s.orderRepo.GetOrderStatistics(startDate, endDate)
	if err != nil {
		return report, err
	}

	report.NoOfOrders = uint(orderCount)
	report.TotalAmount = totalAmount

	last7DaysOrderCount, last7DaysTotalAmount, err := s.orderRepo.GetOrderStatistics(last7DaysStartDate, currentTime)
	if err != nil {
		return report, err
	}

	report.WeeklyOrders = uint(last7DaysOrderCount)
	report.WeeklyTotalAmount = last7DaysTotalAmount

	lastMonthOrderCount, lastMonthTotalAmount, err := s.orderRepo.GetOrderStatistics(lastMonthStartDate, currentTime)
	if err != nil {
		return report, err
	}
	report.MonthlyOrders = uint(lastMonthOrderCount)
	report.MonthlyTotalAmount = lastMonthTotalAmount

	lastYearOrderCount, lastYearTotalAmount, err := s.orderRepo.GetOrderStatistics(lastYearStartDate, currentTime)
	if err != nil {
		return report, err
	}
	report.YearlyOrders = uint(lastYearOrderCount)
	report.YearlyTotalAmount = lastYearTotalAmount

	return report, nil
}

func (s *adminService) UpdatingOrderStatus(orderId uint) (*entity.Orders, error) {

	order, err := s.orderRepo.FindOrderItemById(orderId)
	if err != nil {
		return nil, err
	}
	order.Status = "Delivered"
	_, err = s.orderRepo.UpdateOrderItem(order)
	if err != nil {
		return nil, err
	}
	return nil, nil

}

func (s *adminService) UpdatingReturnOrderStatus(orderId uint) error {

	order, err := s.orderRepo.FindOrderItemById(orderId)
	if err != nil {
		return err
	}
	if order.Status != "Returned" {
		return errors.New("can't update")
	}
	order.Status = "Returned_Recived"
	orderUpdate, err := s.orderRepo.UpdateOrderItem(order)
	if err != nil {
		return err
	}
	orderID := orderUpdate.OrderID

	getOrder, err := s.orderRepo.FindOrderById(orderID)
	userId := getOrder.UserID

	err = s.walletRepo.DepositToWallet(userId, orderUpdate.Price)
	return nil

}

func (s *adminService) AddCoupon(coupon entity.Coupon) error {

	err := s.couponRepo.CreateCoupon(&coupon)
	if err != nil {
		return err
	}
	return nil
}

// func (s *adminService) BlockUser(user *entity.User) error {

// 	user.Block = true
// 	err := s.adminRepo.Block(user)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *adminService) UnblockUser(user *entity.User) error {

// 	user.Block = false
// 	err := s.adminRepo.Unblock(user)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *adminService) Signup(admin *entity.Admin) (*entity.Admin, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}

// 	admin.Password = string(hashedPassword)

// 	err = s.adminRepo.Create(admin)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return admin, nil
// }

// func (s *adminService) Login(admin *entity.Admin) (*entity.Admin, error) {
// 	existingAdmin, err := s.adminRepo.GetByEmail(admin.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(existingAdmin.Password), []byte(admin.Password))
// 	if err != nil {
// 		return nil, errors.New("invalid credentials")
// 	}

// 	return existingAdmin, nil
// }
