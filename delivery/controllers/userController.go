package controllers

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/delivery/helpers"
	"70_Off/domain/auth"
	"70_Off/domain/entity"
	"70_Off/domain/services"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}

}

// UserSignUp godoc
//
//	@Summary		User SignUp
//	@Description	Add new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		deliverymodels.User	true	"User Data"
//	@Success		200		{object}	deliverymodels.User
//	@Router			/userSignup [post]
func (uc *UserController) UserSignUp(c *gin.Context) {

	user := deliverymodels.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := uc.userService.CreateUsers(&user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"key": key, "message": "Verify phoneNumber,Otp Sent Sucessfully"})
}

// UserSignUpVerification godoc
//
//	@Summary		User Phone Number Verification
//	@Description	User SignUp Phone Number Verification through OTP
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.UserLoginWithOtp	true	"User Data"
//	@Success		200		{object}	entity.UserLoginWithOtp
//	@Router			/userSignupVerification [post]
func (uc *UserController) UserSignUpVerification(c *gin.Context) {
	var userOtpLoginRequest entity.UserLoginWithOtp

	if err := c.ShouldBind(&userOtpLoginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userOtpLoginRequest.ResendOtp == true {
		key, err1 := uc.userService.UserOtpSendRequst(&userOtpLoginRequest)
		if err1 != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"key": key, "message": "OTP resent successfully"})

	}

	_, err1 := uc.userService.UserOtpValidation(&userOtpLoginRequest)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User verified successfully"})
}

// UserLogin godoc
//
//	@Summary		User Can Login
//	@Description	User Can Login With UserName And Password
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		deliverymodels.UserLoginInputs	true	"User Data"
//	@Success		200		{object}	deliverymodels.UserLoginInputs
//	@Router			/userLogin [post]
func (uc *UserController) UserLogin(c *gin.Context) {
	var userLoginRequest deliverymodels.UserLoginInputs

	if err := c.ShouldBind(&userLoginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err1 := uc.userService.UserValidateLogin(&userLoginRequest)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}
	token, err := auth.GenerateJWT(user.Email, user.ID, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "70off.online", true, false)

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}

// OTP_Login
//
//	@Summary		User Login With OTP
//	@Description	User Can Login with OTP Without Password
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.UserLoginWithOtp	true	"User Data"
//	@Success		200		{object}	entity.UserLoginWithOtp
//	@Router			/userLoginWithOtp [post]
func (uc *UserController) UserLoginWithOtp(c *gin.Context) {
	var userOtpLoginRequest entity.UserLoginWithOtp

	if err := c.BindJSON(&userOtpLoginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err1 := uc.userService.UserOtpSendRequst(&userOtpLoginRequest)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	response := gin.H{
		"key":     key,
		"message": "otp Send sucessfully",
	}

	c.JSON(http.StatusOK, response)
}

// UserLoginVerification
//
//	@Summary		User Phone Number Verification
//	@Description	User Login  Verification through OTP
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.UserLoginWithOtp	true	"User Data"
//	@Success		200		{object}	entity.UserLoginWithOtp
//	@Router			/userLoginWithOtpValidation [post]
func (uc *UserController) UserLoginWithOtpValidation(c *gin.Context) {
	var userOtpLoginRequest entity.UserLoginWithOtp

	if err := c.BindJSON(&userOtpLoginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userOtpLoginRequest.ResendOtp == true {
		key, err1 := uc.userService.UserOtpSendRequst(&userOtpLoginRequest)
		if err1 != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"key": key, "message": "OTP resent successfully"})

	}

	user, err1 := uc.userService.UserOtpValidation(&userOtpLoginRequest)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	token, err := auth.GenerateJWT(user.Email, user.ID, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "70off.online", true, false)

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
}

// AddAddress godoc
//
//	@Summary		Add Address For User
//	@Description	User Can Add Address After SignUp
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		deliverymodels.Address	true	"User Data (required)"
//	@Success		200		{object}	deliverymodels.Address
//	@Router			/v1/addAddress [post]
func (uc *UserController) AddAddress(c *gin.Context) {
	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var address deliverymodels.Address
	if err := c.BindJSON(&address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err1 := uc.userService.AddUserAddress(&address, uint(userID))
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Address updated successfully",
	})
}

// AddToCart godoc
//
//	@Summary		Add Cart Items
//	@Description	User Can Add Product to Cart
//	@Tags			user
//
//	@param			productID	path		string	true	"ProductID(required)"
//	@param			quantity	path		string	true	"quantity(required)"
//	@param			sizeID		path		string	true	"sizeID(required)"
//
//	@Success		200			{string}	string	"Success Message"
//	@Router			/v1/addCart/{productID}/{quantity}/{sizeID} [post]
func (uc *UserController) AddCartItem(c *gin.Context) {

	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		fmt.Println("\n", productID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid productID"})
		return
	}

	sizeID, err := strconv.Atoi(c.Param("sizeID"))
	if err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid sizeID"})
		return
	}

	quantity, err := strconv.Atoi(c.Param("quantity"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}
	err1 := uc.userService.AddItemToCart(uint(productID), uint(userID), uint(quantity), uint(sizeID))
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart item added successfully"})

}

// Cart
//
//	@Summary		Display User Cart
//	@Description	Display Products in Cart
//	@Tags			user
//	@Produce		json
//	@Success		200	{object}	deliverymodels.ViewCartDetail
//	@Router			/v1/diplayCart [get]
func (uc *UserController) DisplayCart(c *gin.Context) {
	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	cartItems, totalAmount, err := uc.userService.DisplayCart(uint(userID))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"message":     "Cart",
		"cartItems":   cartItems,
		"totalAmount": totalAmount,
	}

	c.JSON(http.StatusOK, response)

}

// RemoveCart godoc
//
//	@Summary		Remove Products From Cart
//	@Description	User Can Remove Product From Cart
//	@Tags			user
//	@param			productItemID	path		string	true	"ProductID(required)"
//	@param			quantity		path		string	true	"quantity(required)"
//	@Success		200				{string}	string	"Success Message"
//	@Router			/v1/removeCartItem/{productItemID}/{quantity} [put]
func (uc *UserController) RemoveCartItem(c *gin.Context) {

	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	productItemID, err := strconv.Atoi(c.Param("productItemID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid productItemID"})
		return
	}

	quantity, err := strconv.Atoi(c.Param("quantity"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid cartItemID"})
		return
	}

	err = uc.userService.RemoveCartItem(uint(userID), uint(productItemID), uint(quantity))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item removed successfully"})
}

// CheckOut
//
//	@Summary		Check Out Cart
//	@Description	User Can CheckOut Products From Cart
//	@Tags			user
//	@Produce		json
//	@Success		200	{object}	deliverymodels.ViewCartDetail
//	@Router			/v1/orderProduct [get]
func (uc *UserController) CheckOutCart(c *gin.Context) {

	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	cartItems, adresses, totalAmount, wallet, coupons, err1 := uc.userService.CheckOutCartProcess(uint(userID))
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	response := gin.H{
		"message":           "Choose your Address",
		"addresses":         adresses,
		"cartItems":         cartItems,
		"totalAmount":       totalAmount,
		"wallet balance":    wallet,
		"Available Coupons": coupons,
	}

	c.JSON(http.StatusOK, response)
}

// Order Confirmation
//
//	@Summary		Order Confirmation
//	@Description	User Can Choose Address and Payment
//	@Tags			user
//	@Param			coupon			query		string								false	"coupon"
//	@Param			wallet			query		string								false	"Wallet"
//	@Param			paymentTypeID	path		string								true	"Payment(required)"
//	@Param			addressID		path		string								true	"Address(required)"
//	@Success		200				{object}	[]deliverymodels.ViewOrdersForUsers	"Success Message"
//	@Router			/v1/CheckOutCartProcess/{paymentTypeID}/{addressID}/{coupon} [post]
func (uc *UserController) PaymentSelection(c *gin.Context) {
	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	wallet := c.DefaultQuery("wallet", "0")
	walletAmt, err := strconv.Atoi(wallet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet"})
		return
	}

	coupon := c.DefaultQuery("coupon", "")
	// coupon, err := strconv.Atoi(couponID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coupon"})
	// 	return
	// }

	paymentTypeID, err := strconv.Atoi(c.Param("paymentTypeID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Payment Type"})
		return
	}

	addressID, err := strconv.Atoi(c.Param("addressID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid addressID"})
		return
	}

	session := sessions.Default(c)
	session.Set("addressID", addressID)
	session.Set("wallet", walletAmt)
	session.Set("coupon", coupon)
	session.Save()

	if paymentTypeID == 1 {
		order, invoice, err := uc.userService.CheckOutCart(uint(addressID), uint(userID), 2, coupon, float64(walletAmt))
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		response := gin.H{
			"message": "Order placed",
			"invoice": invoice,
			"order":   order,
		}

		c.JSON(http.StatusOK, response)
	} else {
		PayableAmount, err := uc.userService.RazorpayCheckOut(uint(addressID), uint(userID),
			uint(paymentTypeID), coupon, float64(walletAmt))

		razorpayid, err := helpers.GenerateRazorpayOrder(uint(PayableAmount) * 100)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		response := gin.H{
			"message":    "payment",
			"razorpayid": razorpayid,
		}

		c.JSON(http.StatusOK, response)
	}

}

// Razor Pay
//
//	@Summary		Razor Pay
//	@Description	User Can Pay Through RazorPay
//	@Tags			user
//	@Param			sign	path		string	true	"Sign(required)"
//	@Param			razorid	path		string	true	"Razorid(required)"
//	@Param			payid	path		string	true	"Payid(required)"
//	@Success		200		{object}	deliverymodels.ViewInvoiceForUsers
//	@Router			/v1/razorpayVerification/{sign}/{razorid}/{payid} [post]
func (uc *UserController) RazorPayVerification(c *gin.Context) {
	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	session := sessions.Default(c)
	addressID := session.Get("addressID")
	wallet := session.Get("wallet")
	couponCode := session.Get("coupon")
	Signature := c.Param("sign")
	razorId := c.Param("razorid")
	paymentId := c.Param("payid")

	walletAmt, ok := wallet.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet amount"})
		return
	}

	coupon, ok := couponCode.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coupon"})
		return
	}

	address, ok := addressID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}

	if err := uc.userService.ExecuteRazorPaymentVerification(Signature, razorId, paymentId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"massage": "Payment successful"})

	order, invoice, err := uc.userService.CheckOutCart(uint(address), uint(userID), 2, coupon, float64(walletAmt))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	session.Delete("addressID")
	session.Delete("wallet")
	session.Delete("couponID")
	session.Save()

	response := gin.H{
		"message": "Order placed",
		"invoice": invoice,
		"order":   order,
	}

	c.JSON(http.StatusOK, response)
}

// Order History
//
//	@Summary		order history
//	@Description	order history,
//	@Tags			user
//	@param			page	query		string	false	"Page"
//	@param			limit	query		string	false	"Limit"
//	@Success		200		{object}	deliverymodels.ViewOrdersForUsers
//	@Router			/v1/orderHistory [get]
func (uc *UserController) OrderHistory(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 || limitNum > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	offset := (pageNum - 1) * limitNum

	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err5 := strconv.Atoi(userIDCookie)
	if err5 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	order, err1 := uc.userService.OrderHistory(uint(userID), offset, limitNum)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	response := gin.H{
		"message": "Order History",
		"orders":  order,
	}

	c.JSON(http.StatusOK, response)

}

// Cancel  Order
//
//	@Summary		Cancel  Order By User
//	@Description	User Can Cancel Order Using ID,
//	@Tags			user
//	@param			orderID	path		string	true	"orderID"
//	@Success		200		{string}	string	"Success Message"
//	@Router			/v1/cancelOrderByUser/{orderID} [put]
func (uc *UserController) CancelOrderByUser(c *gin.Context) {

	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid orderID"})
		return
	}

	err1 := uc.userService.CancelOrderByUser(uint(orderID))
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order Canceled"})

}

// Reset Password
//
//	@Summary		Reset Password - OTP
//	@Description	Reset Password Using OTP Verification
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.UserLoginWithOtp	true	"User Data"
//	@Success		200		{string}	message
//	@Router			/resetPasswordRequst [get]
func (uc *UserController) ResetPasswordRequestUser(c *gin.Context) {

	user := &entity.UserLoginWithOtp{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := uc.userService.UserOtpSendRequst(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": key, "message": "OTP sent successfully,verify OTP"})

}

// Reset Password Verification
//
//	@Summary		Phone Number Verification
//	@Description	User PhoneNumber Verification For Reset Password
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.UserLoginWithOtp	true	"User Data"
//	@Success		200		{string}	message
//	@Router			/resetPassword [put]
func (uc *UserController) ResetPasswordUser(c *gin.Context) {

	var userOtpLoginRequest entity.UserLoginWithOtp

	if err := c.BindJSON(&userOtpLoginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userOtpLoginRequest.ResendOtp == true {
		key, err1 := uc.userService.UserOtpSendRequst(&userOtpLoginRequest)
		if err1 != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"key": key, "message": "OTP sent successfully,verify OTP"})

	}

	_, err1 := uc.userService.UserOtpValidation(&userOtpLoginRequest)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	hashedPassword, err4 := bcrypt.GenerateFromPassword([]byte(userOtpLoginRequest.Password), bcrypt.DefaultCost)
	if err4 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err4.Error()})
		return
	}
	user := deliverymodels.UpdateUser{
		Password: string(hashedPassword),
	}

	getuser, err := uc.userService.GetByPhoneNumber(userOtpLoginRequest.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	err = uc.userService.UserUpdate(getuser.ID, &user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password Updated Sucessfully"})

}

// Profile Updates
//
//	@Summary		User Profile Updates
//	@Description	User Can Updated User Profile Details
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		deliverymodels.UpdateUser	true	"User Data"
//	@Success		200		{string}	message
//	@Router			/v1/updateUserDetails [put]
func (uc *UserController) UserUpdates(c *gin.Context) {

	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user deliverymodels.UpdateUser

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := uc.userService.UserUpdate(uint(userID), &user)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error in update user": err1.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Details Updated Sucessfully"})

}

// Search
//
//	@Summary		Search Product
//	@Description	Search Product By Name,
//	@Tags			user
//	@param			search	path		string	false	"Search"
//	@param			page	query		string	false	"Page"
//	@param			limit	query		string	false	"Limit"
//	@Success		200		{object}	deliverymodels.ViewProductDetail
//	@Router			/v1/searchByName/{search} [get]
func (uc *UserController) SearchProductsByName(c *gin.Context) {
	input := c.Param("search")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	products, err := uc.userService.SearchProductsUsingName(input, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Results",
		"filtered": products,
	})
}

// Filter
//
//	@Summary		Filter Product
//	@Description	Filter Product By Brand,
//	@Tags			user
//	@param			brandID	path		string	false	"BrandID"
//	@param			page	query		string	false	"Page"
//	@param			limit	query		string	false	"Limit"
//	@Success		200		{object}	deliverymodels.ViewProductDetail
//	@Router			/v1/filterByBrand/{brandID} [get]
func (uc *UserController) ProductCategoryFiltering(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	brandID, err := strconv.Atoi(c.Param("brandID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid brandID"})
		return
	}

	var brand string
	if brandID == 1 {
		brand = "SCOTT"
	} else {
		brand = "BITWIN"
	}
	product, err := uc.userService.ProductCategoryFiltering(uint(brandID), page, limit)
	if err != nil {
		return
	}
	productDisplay := make([]deliverymodels.ViewProductDetail, len(product))
	for n, prodct := range product {
		productDetails, err := uc.userService.GetProductDetailsByProductId(prodct.ID)
		if err != nil {
			return
		}
		productDisplay[n].Price = productDetails.Price
		productDisplay[n].Brand = brand
		productDisplay[n].Description = prodct.Description
		productDisplay[n].Name = prodct.Name
		// productDisplay[n].Sizes = make(string)["S","M","L"]
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  brand,
		"filtered": productDisplay,
	})

}

// Return  Order
//
//	@Summary		Return  Order By User
//	@Description	User Can Return Order Using ID,
//	@Tags			user
//	@param			orderID	path		string	true	"OrderID"
//	@Success		200		{string}	string	"Success Message"
//	@Router			/v1/returnOrderByUser/{orderID} [put]
func (uc *UserController) ReturnOrderByUser(c *gin.Context) {
	orderItemID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid orderID"})
		return
	}

	err = uc.userService.OrderReturn(uint(orderItemID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't return"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "return requst submit successfully"})

}

// Wallet
//
//	@Summary		Wallet Balance
//	@Description	User Can Check Wallet Balance,
//	@Tags			user
//	@Success		200	{string}	string	"Success Message"
//	@Router			/v1/walletBalance [get]
func (uc *UserController) DisplayWalletBalance(c *gin.Context) {
	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	balance, err := uc.userService.WalletBalance(uint(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't fetch balance"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "70% Off Wallet Balance",
		"Balance": balance,
	})

}

// Product  Details
//
//	@Summary		Product Details
//	@Description	User Can View Product Details,
//	@Tags			user
//	@param			productID	path		string	true	"ProductID"
//	@Success		200			{string}	string	"Success Message"
//	@Router			/v1/productDetail/{productID} [get]
func (uc *UserController) ProductDetails(c *gin.Context) {

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid productID"})
		return
	}
	productDetail, err := uc.userService.ProductDetail(uint(productID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't fetch Product Details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Product Detail": productDetail,
	})
}

// User Profile
//
//	@Summary		User Profile
//	@Description	To Display User Details And Address,
//	@Tags			user
//	@Success		200	{object}	deliverymodels.ViewUserProfile
//	@Router			/v1/userProfile [get]
func (uc *UserController) UserProfile(c *gin.Context) {
	userIDCookie, err := c.Cookie("userID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := strconv.Atoi(userIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, address, err := uc.userService.UserProfile(uint(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Cant't Fetch Details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User Deatil": user,
		"Address":     address,
	})

}

// LogOut
//
//	@Summary		Logout
//	@Description	User Logout,
//	@Tags			user
//	@Success		200	{string}	string	"Success Message"
//	@Router			/v1/logOut [post]
func (uc *UserController) LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "70off.online", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
