package services

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/delivery/helpers"
	"70_Off/domain/entity"
	"70_Off/domain/repository"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/razorpay/razorpay-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUsers(users *deliverymodels.User) (string, error)
	AddUserAddress(address *deliverymodels.Address, userId uint) error
	UserValidateLogin(userLoginRequest *deliverymodels.UserLoginInputs) (*entity.User, error)
	UserOtpSendRequst(UserLoginRequst *entity.UserLoginWithOtp) (string, error)
	UserOtpValidation(UserLoginRequst *entity.UserLoginWithOtp) (*entity.User, error)
	UserProfile(userId uint) (*deliverymodels.ViewUserProfile, *[]deliverymodels.ViewAddressForUser, error)

	AddItemToCart(productSlug string, userId, quantity, sizeId uint) error
	RemoveCartItem(userId, productItemId, quantity uint) error
	CheckOutCartProcess(userId uint) (*[]deliverymodels.ViewCartDetail, *[]deliverymodels.ViewAddressForUser,
		float64, float64, *[]deliverymodels.Coupon, error)
	DisplayCart(userId uint) (*[]deliverymodels.ViewCartDetail, float64, error)
	ExecutePurchaseRazorPay(cartId, addressId uint) (string, uint, error)
	ExecuteRazorPaymentVerification(Signature, razorId, paymentId string) error
	CheckOutCart(addressId, userId, paymentTypeID uint, coupon string, walleAmt float64) (*[]deliverymodels.ViewOrdersForUsers,
		*deliverymodels.ViewInvoiceForUsers, error)

	RazorpayCheckOut(addressId, userId, paymentTypeID uint, coupon string, walleAmt float64) (float64, error)
	CancelOrderByUser(orderId uint) error
	OrderHistory(userId uint, offset, limit int) (*[]deliverymodels.ViewOrdersForUsers, error)
	GetByPhoneNumber(phoneNumber string) (*entity.User, error)
	UserUpdate(userID uint, user *deliverymodels.UpdateUser) error
	GetByUserId(userId uint) (*entity.User, error)
	GetCartByUserId(userId uint) (*entity.Cart, error)
	ProductCategoryFiltering(brandId uint, page, limit int) ([]entity.Product, error)
	GetProductDetailsByProductId(id uint) (*entity.ProductDetails, error)
	SearchProductsUsingName(input string, page, limit int) ([]*deliverymodels.ViewProductDetail, error)
	OrderReturn(orderItemId uint) error
	WalletBalance(userId uint) (float64, error)
	ProductDetail(producId uint) (*deliverymodels.ViewProductDetail, error)
}

type userService struct {
	userRepo          repository.UserRepository
	cartRepo          repository.CartRepository
	cartItemRepo      repository.CartItemRepository
	orderRepo         repository.OrderRepository
	invoiceRepo       repository.InvoiceRepository
	otpValidationRepo repository.OtpValidationRepository
	inventoryRepo     repository.Inventory
	productRepo       repository.Product
	paymentRepo       repository.PaymentRepository
	walletRepo        repository.WalletRepository
	couponRepo        repository.CouponRepository
}

func NewUserService(userRepo repository.UserRepository, cartRepo repository.CartRepository,
	cartItemRepo repository.CartItemRepository, orderRepo repository.OrderRepository,
	invoiceRepo repository.InvoiceRepository, otpValidationRepo repository.OtpValidationRepository,
	inventoryRepo repository.Inventory, productRepo repository.Product,
	paymentRepo repository.PaymentRepository, walletRepo repository.WalletRepository,
	couponRepo repository.CouponRepository) UserService {

	return &userService{userRepo: userRepo, cartRepo: cartRepo, cartItemRepo: cartItemRepo,
		orderRepo: orderRepo, invoiceRepo: invoiceRepo, otpValidationRepo: otpValidationRepo,
		inventoryRepo: inventoryRepo, productRepo: productRepo, paymentRepo: paymentRepo,
		walletRepo: walletRepo, couponRepo: couponRepo}
}

func (s *userService) UserUpdate(userID uint, user *deliverymodels.UpdateUser) error {
	toUpdateUser, err2 := s.GetByUserId(userID)
	if err2 != nil {
		return err2
	}
	if user.FirstName != "" {
		toUpdateUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		toUpdateUser.LastName = user.LastName
	}
	if user.Email != "" {
		toUpdateUser.Email = user.Email
	}
	if user.Username != "" {
		toUpdateUser.Username = user.Username
	}

	err := s.userRepo.Update(toUpdateUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetByPhoneNumber(phoneNumber string) (*entity.User, error) {

	user, err := s.userRepo.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByUserId(userId uint) (*entity.User, error) {

	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUsers(user *deliverymodels.User) (string, error) {
	email := user.Email
	username := user.Username
	phoneNumber := user.PhoneNumber

	if _, err := s.userRepo.GetByEmail(email); err == nil {
		return "", errors.New("email already exists")
	}

	if _, err := s.userRepo.GetUserByUserName(username); err == nil {
		return "", errors.New("username already exists")
	}

	if _, err := s.userRepo.GetByPhoneNumber(phoneNumber); err == nil {
		return "", errors.New("phoneNumber already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	toCreateUser := entity.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    string(hashedPassword),
		Username:    user.Username,
	}

	if err := s.userRepo.Create(&toCreateUser); err != nil {
		return "", err
	}
	key, err2 := helpers.SendOtp(user.PhoneNumber)
	if err2 != nil {
		return "", errors.New("otp sending failed")
	}
	otpdata := entity.OtpValidation{
		Key: key, PhoneNumber: phoneNumber,
	}
	otpdataCreated, err3 := s.otpValidationRepo.Create(&otpdata)
	if err3 != nil {
		return "", errors.New("otp sending failed")
	}

	return otpdataCreated.Key, nil
}

func (s *userService) AddUserAddress(address *deliverymodels.Address, userId uint) error {

	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return err
	}
	forCreateAddress := entity.Address{
		HouseNo:    address.HouseNo,
		Street:     address.Street,
		City:       address.City,
		State:      address.State,
		PostalCode: address.PostalCode,
		UserID:     userId,
		User:       *user,
	}

	err = s.userRepo.AddAddress(&forCreateAddress)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UserValidateLogin(userLoginRequest *deliverymodels.UserLoginInputs) (*entity.User, error) {
	username := userLoginRequest.Username
	password := userLoginRequest.Password

	user, err := s.userRepo.GetUserByUserName(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user.Block {
		return nil, errors.New("user Blocked")
	}

	if !user.Verified {
		return nil, errors.New("can't login, User need to verify")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *userService) UserOtpSendRequst(userLoginRequest *entity.UserLoginWithOtp) (string, error) {

	phoneNumber := userLoginRequest.PhoneNumber
	user, err1 := s.userRepo.GetByPhoneNumber(phoneNumber)
	if err1 != nil {
		return "", errors.New("invalid credentials")
	}

	if user.Block {
		return "", errors.New("user Blocked")
	}

	key, err2 := helpers.SendOtp(user.PhoneNumber)
	if err2 != nil {
		return "", errors.New("otp sending failed")
	}
	otpdata := entity.OtpValidation{
		Key: key, PhoneNumber: phoneNumber,
	}
	otpdataCreated, err3 := s.otpValidationRepo.Create(&otpdata)
	if err3 != nil {
		return "", errors.New("otp sending failed")
	}
	return otpdataCreated.Key, nil

}

func (s *userService) UserOtpValidation(userLoginRequest *entity.UserLoginWithOtp) (*entity.User, error) {
	userOtpData, err1 := s.otpValidationRepo.FindByKey(userLoginRequest.Key)
	if err1 != nil {
		return nil, errors.New("invalid key")
	}
	phoneNumber := "+91" + userOtpData.PhoneNumber
	otp := userLoginRequest.Otp

	err2 := helpers.CheckOtp(phoneNumber, otp)
	if err2 != nil {
		return nil, errors.New("otp validation failed")
	}
	user, err := s.userRepo.GetByPhoneNumber(userOtpData.PhoneNumber)
	if err != nil {
		return nil, errors.New("user PhoneNumber Does'nt match with dataBase")
	}
	if !user.Verified {
		user.Verified = true
		err3 := s.UpdateUser(user)
		if err3 != nil {
			return nil, errors.New("user verified update failed")
		}
	}
	return user, nil

}

func (s *userService) AddItemToCart(productSlug string, userId, quantity, sizeId uint) error {
	product, err := s.productRepo.ProductGetBySlug(productSlug)
	if err != nil {
		return err
	}
	fmt.Println("\nProductGetBySlug\n", product)
	productDetails, err := s.productRepo.ProductItemGetByProductIdAndSizeId(product.ID, sizeId)
	if err != nil {
		return err
	}

	cart, _ := s.cartRepo.FindCartByUserID(userId)
	if cart == nil {
		createCart := entity.Cart{
			UserID:          userId,
			TotalCartAmount: 0,
		}
		var err1 error
		cart, err1 = s.cartRepo.SaveCart(&createCart)
		if err1 != nil {
			return err1
		}
	}

	cartItems, err := s.cartItemRepo.FindAllCartItemsByCartID(cart.ID)
	if err != nil {
		return err
	}

	found := false
	for _, cartItem := range cartItems {
		if cartItem.ProductDetailsID == productDetails.ID {
			found = true
			cartItem.Quantity = cartItem.Quantity + quantity
			cartItem.Price = cartItem.Price + (productDetails.Price * float64(quantity))
			_, err := s.cartItemRepo.UpdateCartItem(&cartItem)
			if err != nil {
				return errors.New("can't update cart item")
			}
			cart.TotalCartAmount += (productDetails.Price * float64(quantity))
			_, err = s.cartRepo.UpdateCart(cart)
			if err != nil {
				return err
			}

			break
		}
	}

	if found {
		return nil
	}
	createCartItem := &entity.CartItem{
		CartID:           cart.ID,
		ProductDetailsID: productDetails.ID,
		Quantity:         quantity,
		Price:            float64(quantity) * productDetails.Price,
	}

	_, err = s.cartItemRepo.SaveCartItem(createCartItem)
	if err != nil {
		return err
	}

	cart.TotalCartAmount += createCartItem.Price
	_, err = s.cartRepo.UpdateCart(cart)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) RemoveCartItem(userId, productId, quantity uint) error {

	cart, err := s.cartRepo.FindCartByUserID(userId)
	if err != nil {
		return err
	}

	cartItem, err := s.cartItemRepo.FindCartItemByUserIdAndProductDetailsID(userId, productId)
	if err != nil {
		return err
	}

	if err = s.cartItemRepo.DeleteCartItemWithQuantity(cartItem.ID, quantity); err != nil {
		return err
	}

	productItemPrice := cartItem.Price / float64(cartItem.Quantity)

	cart.TotalCartAmount -= productItemPrice * float64(quantity)

	_, err1 := s.cartRepo.UpdateCart(cart)
	if err1 != nil {
		return err1
	}
	return nil
}

func (s *userService) UpdateUser(user *entity.User) error {
	err := s.userRepo.Update(user)
	if err != nil {
		return err
	}
	return nil
}
func (s *userService) DisplayCart(userId uint) (*[]deliverymodels.ViewCartDetail, float64, error) {
	cart, err := s.cartRepo.FindCartByUserID(userId)
	if err != nil {
		return nil, 0, err
	}
	cartItems, err := s.cartItemRepo.FindAllCartItemsByCartIDForUser(cart.ID)
	if err != nil {
		return nil, 0, err
	}

	totalAmount := cart.TotalCartAmount

	return &cartItems, totalAmount, nil
}

func (s *userService) CheckOutCartProcess(userId uint) (*[]deliverymodels.ViewCartDetail,
	*[]deliverymodels.ViewAddressForUser, float64, float64, *[]deliverymodels.Coupon, error) {
	cart, err := s.cartRepo.FindCartByUserID(userId)
	if err != nil {
		return nil, nil, 0, 0, nil, err
	}

	cartItems, err := s.cartItemRepo.FindAllCartItemsByCartIDForUser(cart.ID)
	if err != nil {
		return nil, nil, 0, 0, nil, err
	}

	totalAmount := cart.TotalCartAmount

	addresses, err := s.userRepo.GetAllAddressesByUserId(userId)
	if err != nil {
		return nil, nil, 0, 0, nil, err
	}

	toDisplayAddresses := make([]deliverymodels.ViewAddressForUser, len(*addresses))

	for count, address := range *addresses {
		toDisplayAddress := deliverymodels.ViewAddressForUser{
			FirstName:   address.User.FirstName,
			LastName:    address.User.LastName,
			PhoneNumber: address.User.PhoneNumber,
			City:        address.City,
			HouseNo:     address.HouseNo,
			Street:      address.Street,
			State:       address.State,
			PostalCode:  address.PostalCode,
		}
		toDisplayAddresses[count] = toDisplayAddress
	}

	getcoupons, err := s.couponRepo.GetAllCoupons()
	if err != nil {
		return &cartItems, &toDisplayAddresses, totalAmount, 0, nil, fmt.Errorf("failed to get coupons: %w", err)
	}

	coupons := make([]deliverymodels.Coupon, len(getcoupons))
	for i, coupon := range getcoupons {
		copier.Copy(&coupons[i], coupon)
	}

	wallet, err := s.walletRepo.GetWalletBalance(userId)
	if err != nil {
		return &cartItems, &toDisplayAddresses, totalAmount, 0, &coupons, nil

	}

	return &cartItems, &toDisplayAddresses, totalAmount, wallet, &coupons, nil
}

func (s *userService) CheckOutCart(addressID, userID, paymentTypeID uint, coupon string, walletAmt float64) (
	*[]deliverymodels.ViewOrdersForUsers, *deliverymodels.ViewInvoiceForUsers, error) {

	if walletAmt > 0 {
		err := s.walletRepo.WithdrawFromWallet(userID, walletAmt)
		if err != nil {
			return nil, nil, errors.New("can't withdrow from wallet")
		}

	}

	_, err := s.userRepo.GetAddressById(addressID)
	if err != nil {
		return nil, nil, err
	}

	cart, err := s.cartRepo.FindCartByUserID(userID)
	if err != nil {
		return nil, nil, err
	}

	cartItems, err := s.cartItemRepo.FindAllCartItemsByCartID(cart.ID)
	if err != nil {
		return nil, nil, err
	}
	if cartItems == nil {
		return nil, nil, errors.New("empty cart")
	}

	var getCoupon = &entity.Coupon{}
	// discount := 0.00
	if coupon != "" {
		getCoupon, err = s.couponRepo.GetCouponByCouponCode(coupon)
		if err != nil {
			return nil, nil, errors.New("coupon not valid")
		}
	}

	totalAmount := cart.TotalCartAmount
	discount := getCoupon.DiscountPercentage * totalAmount / 100
	if getCoupon.MaximumDiscountAmount < discount {
		discount = getCoupon.MaximumDiscountAmount
	}

	order := &entity.Orders{
		UserID:          userID,
		AddressID:       addressID,
		TotalBillAmount: totalAmount,
		PaymentTypeID:   paymentTypeID,
	}

	createdOrder, err := s.orderRepo.SaveOrder(order)
	if err != nil {
		return nil, nil, err
	}

	var productItems []entity.ProductDetails
	for _, cartItem := range cartItems {
		orderItem := &entity.OrderItems{
			OrderID:          createdOrder.ID,
			ProductDetailsID: cartItem.ProductDetailsID,
			Quantity:         cartItem.Quantity,
			Price:            cartItem.Price,
			Status:           "order confirmed",
		}

		_, err = s.orderRepo.SaveOrderItem(orderItem)
		if err != nil {
			return nil, nil, err
		}

		productItem, err3 := s.productRepo.ProductItemGetByID(cartItem.ProductDetailsID)
		if err3 != nil {
			return nil, nil, err3
		}

		productItem.Quantity = productItem.Quantity - cartItem.Quantity
		productItems = append(productItems, *productItem)

		err2 := s.cartItemRepo.DeleteCartItem(&cartItem)
		if err2 != nil {
			return nil, nil, err2
		}
	}

	payment := entity.Payment{
		OrderId:       createdOrder.ID,
		PaymentTypeId: paymentTypeID,
		PayableAmount: totalAmount - discount - walletAmt,
	}
	if paymentTypeID == 2 {
		payment.PaymentStatus = true
	}

	paymentCreated, err := s.paymentRepo.SavePayment(&payment)
	if err != nil {
		return nil, nil, err
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, nil, err
	}

	invoice := &entity.Invoice{
		Name:           user.FirstName + " " + user.LastName,
		OrderId:        order.ID,
		TotalAmount:    totalAmount,
		PayableAmount:  totalAmount - discount - walletAmt,
		DiscountAmount: discount,
		PaymentID:      paymentCreated.ID,
	}

	invoiceCreated, err := s.invoiceRepo.SaveInvoice(invoice)
	if err != nil {
		return nil, nil, err
	}

	for _, productitem := range productItems {

		err := s.productRepo.ProductItemUpdate(&productitem)
		if err != nil {
			return nil, nil, err
		}
	}

	getcart, err5 := s.cartRepo.FindCartByID(cart.ID)
	if err5 != nil {
		return nil, nil, err5
	}

	err6 := s.cartRepo.DeleteCart(getcart)
	if err6 != nil {
		return nil, nil, err6
	}

	toDisplayOrder, err := s.orderRepo.FindAllOrderItemsByOrderID(createdOrder.ID)
	if err != nil {
		return nil, nil, err
	}

	_, err = s.paymentRepo.FindPaymentById(paymentCreated.ID)
	if err != nil {
		return nil, nil, err
	}

	var paytype string
	if paymentTypeID == 1 {
		paytype = "Cash on Delivery"
	} else {
		paytype = "Razor Pay"
	}

	toDisplayInvoice := deliverymodels.ViewInvoiceForUsers{
		Name:          invoiceCreated.Name,
		OrderId:       invoiceCreated.OrderId,
		TotalAmount:   invoiceCreated.TotalAmount,
		Discount:      discount,
		BillAmount:    invoiceCreated.PayableAmount,
		PaymentStatus: paytype,
		FromWallet:    walletAmt,
	}

	return &toDisplayOrder, &toDisplayInvoice, nil
}

func (s *userService) RazorpayCheckOut(addressId, userId, paymentTypeID uint, coupon string, walleAmt float64) (
	float64, error) {
	if walleAmt > 0 {
		err := s.walletRepo.WithdrawFromWallet(userId, walleAmt)
		if err != nil {
			return 0, errors.New("can't withdrow from wallet")
		}
	} else {
		walleAmt = 0
	}

	_, err := s.userRepo.GetAddressById(addressId)
	if err != nil {
		return 0, err
	}

	cart, err := s.cartRepo.FindCartByUserID(userId)
	if err != nil {
		return 0, err
	}

	cartItems, err := s.cartItemRepo.FindAllCartItemsByCartID(cart.ID)
	if err != nil {
		return 0, err
	}
	if cartItems == nil {
		return 0, errors.New("empty cart")
	}

	totalAmount := cart.TotalCartAmount

	var getCoupon = &entity.Coupon{}
	if coupon != "" {
		getCoupon, err = s.couponRepo.GetCouponByCouponCode(coupon)
		if err != nil {
			return 0, errors.New("coupon not valid")
		}

		discount := getCoupon.DiscountPercentage * totalAmount / 100
		if getCoupon.MaximumDiscountAmount < discount {
			discount = getCoupon.MaximumDiscountAmount
		}
		return totalAmount - discount - walleAmt, nil
	}

	return totalAmount - walleAmt, nil

}

func (s *userService) ExecutePurchaseRazorPay(userId, addressId uint) (string, uint, error) {
	cart, _ := s.cartRepo.FindCartByID(uint(userId))

	client := razorpay.NewClient("rzp_test_O6q2DXJHecJBHI", "MU9PWzkhTBSCkPnxEUOAZdYW")

	data := map[string]interface{}{
		"amount":   int(cart.TotalCartAmount * 100),
		"currency": "INR",
		"receipt":  "101",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", 0, errors.New("payment not initiated")
	}
	razorId, _ := body["id"].(string)

	return razorId, 0, nil
}

func (s *userService) ExecuteRazorPaymentVerification(Signature, razorId, paymentId string) error {

	secret := "BApvy9EJwAaCeQT3ouZsNP5v"
	data := razorId + "|" + paymentId

	h := hmac.New(sha256.New, []byte(secret))

	_, err := h.Write([]byte(data))

	if err != nil {
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(Signature)) != 1 {
		return errors.New("payment failed")
	}
	return nil
}

func (s *userService) OrderHistory(userId uint, offset, limit int) (*[]deliverymodels.ViewOrdersForUsers, error) {

	orderItems, err := s.orderRepo.ListUserOrderItemPaginated(userId, offset, limit)
	if err != nil {
		return nil, err
	}

	return &orderItems, nil
}

func (s *userService) CancelOrderByUser(orderItemId uint) error {

	orderItem, err := s.orderRepo.FindOrderItemById(orderItemId)
	if err != nil {
		return err
	}
	if orderItem.Status != "order placed " && orderItem.Status != "order confirmed" {
		return errors.New("canecelation Time Exeeded")
	}
	order, err := s.orderRepo.FindOrderById(orderItem.OrderID)
	if err != nil {
		return err
	}

	if order.PaymentTypeID == 2 {
		refund := orderItem.Price
		err := s.walletRepo.DepositToWallet(order.UserID, refund)
		if err != nil {
			return err
		}
	}

	orderItem.Status = "canceled"
	_, err1 := s.orderRepo.UpdateOrderItem(orderItem)
	if err1 != nil {
		return err1
	}
	return nil
}

func (s *userService) GetCartByUserId(userId uint) (*entity.Cart, error) {
	cart, err := s.cartRepo.FindCartByUserID(userId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *userService) ProductCategoryFiltering(brandID uint, page int, limit int) ([]entity.Product, error) {
	product, err := s.productRepo.GetProductsByBrandID(brandID, page, limit)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *userService) GetProductDetailsByProductId(id uint) (*entity.ProductDetails, error) {
	productDetails, err := s.productRepo.ProductItemGetByProductIdAndSizeId(id, 2)
	if err != nil {
		return nil, err
	}
	return productDetails, nil
}
func (s *userService) SearchProductsUsingName(input string, page, limit int) ([]*deliverymodels.ViewProductDetail, error) {
	products, err := s.productRepo.SearchProductByName(input, page, limit)
	if err != nil {
		return nil, err
	}

	productDetails := make([]*deliverymodels.ViewProductDetail, len(products))
	for i, product := range products {
		productDetails[i] = &deliverymodels.ViewProductDetail{
			Name:        product.Name,
			Description: product.Description,
		}
		if product.BrandID == 1 {
			productDetails[i].Brand = "Scott"
		} else {
			productDetails[i].Brand = "BitWin"
		}

		getProductDetails, err := s.productRepo.GetProductDetailsByProductID(product.ID)
		if err != nil {
			return nil, err
		}
		productDetails[i].Price = getProductDetails[0].Price
		productDetails[i].Sizes = append(productDetails[i].Sizes, "Small", "Mediume", "Large ")

	}

	return productDetails, nil
}

func (s *userService) OrderReturn(orderItemId uint) error {

	order, err := s.orderRepo.FindOrderItemById(orderItemId)
	if err != nil {
		return err
	}

	if order.Status != "Delivered" {
		return errors.New("please check order id")
	}
	order.Status = "Returned"
	_, err = s.orderRepo.UpdateOrderItem(order)
	if err != nil {
		return err
	}

	return nil

}

func (s *userService) WalletBalance(userId uint) (float64, error) {
	balance, err := s.walletRepo.GetWalletBalance(userId)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (s *userService) ProductDetail(producId uint) (*deliverymodels.ViewProductDetail, error) {

	var brand string
	product, err := s.productRepo.ProductGetByID(producId)
	if err != nil {
		return nil, err
	}
	if product.BrandID == 1 {
		brand = "Scott"
	} else {
		brand = "Bitwin"
	}

	productsItems, err := s.productRepo.GetProductDetailsByProductID(producId)
	if err != nil {
		return nil, err
	}

	productDetail := &deliverymodels.ViewProductDetail{
		Name:        product.Name,
		Description: product.Description,
		Brand:       brand,
		Price:       productsItems[0].Price,
	}

	productDetail.Sizes = append(productDetail.Sizes, "Small", "Mediume", "Large ")

	return productDetail, nil
}

func (s *userService) UserProfile(userId uint) (*deliverymodels.ViewUserProfile, *[]deliverymodels.ViewAddressForUser, error) {

	addresses, err := s.userRepo.GetAllAddressesByUserId(userId)
	if err != nil {
		return nil, nil, err
	}

	toDisplayAddresses := make([]deliverymodels.ViewAddressForUser, len(*addresses))

	for count, address := range *addresses {
		toDisplayAddress := deliverymodels.ViewAddressForUser{
			FirstName:   address.User.FirstName,
			LastName:    address.User.LastName,
			PhoneNumber: address.User.PhoneNumber,
			City:        address.City,
			HouseNo:     address.HouseNo,
			Street:      address.Street,
			State:       address.State,
			PostalCode:  address.PostalCode,
		}
		toDisplayAddresses[count] = toDisplayAddress
	}

	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return nil, nil, err
	}

	viewUser := deliverymodels.ViewUserProfile{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Username:    user.Username,
	}

	return &viewUser, &toDisplayAddresses, nil
}
