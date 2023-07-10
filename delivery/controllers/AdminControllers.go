package controllers

import (
	"70_Off/delivery/auth"
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/entity"
	"70_Off/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AdminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService services.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}

// AdminSignUp
//
//	@Summary		Admin SignUp
//	@Description	Add new Admin
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			admin	body	deliverymodels.AdminSignupInputs	true	"User Data"
//	@Router			/adminSignup [post]
func (ac *AdminController) AdminSignUp(c *gin.Context) {

	var admin deliverymodels.AdminSignupInputs
	if err := c.BindJSON(&admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.adminService.Create(&admin); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, admin)
}

// AdminLogin
//
//	@Summary		Admin Login
//	@Description	Admin Can Login With UserName And Password
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			admin	body	deliverymodels.AdmionLoginInputs	true	"Admin Data"
//	@Success		200
//	@Router			/adminLogin [post]
func (ac *AdminController) AdminLogin(c *gin.Context) {

	var adminLoginInput deliverymodels.AdmionLoginInputs

	if err := c.BindJSON(&adminLoginInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err1 := ac.adminService.AdminValidateLogin(&adminLoginInput)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error admin validate login": err1.Error()})
		return
	}
	token, err := auth.GenerateJWT(admin.Email, 0, admin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "70off.online", true, false)

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})

}

// Block User
//
//	@Summary		User BLock Toggle
//	@Description	Admin Can Block or Unblock User
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		string	true	"UserID "
//	@Success		200		{string}	message
//	@Router			/v1/admin/userBlockToggle/{userID} [put]
func (ac *AdminController) UserBLockToggle(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid productID"})
		return
	}

	err1 := ac.adminService.UserBlockToggle(uint(userID))
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Change save successfully"})

}

// Report
//
//	@Summary		Order Report
//	@Description	It shows weekly,monthly,yearly and last n days order report
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			days	path		string	true	"Days"
//	@Success		200		{string}	message
//	@Router			/v1/admin/saleReport/{days} [get]
func (ac *AdminController) OrderReport(c *gin.Context) {
	days, err := strconv.Atoi(c.Param("days"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	report, err := ac.adminService.OrderReport(uint(days))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":                                  "Reports",
		"Your Given Count, OrdersReports":          report.NoOfOrders,
		"Your Given Count, TotalAmount For orders": report.TotalAmount,
		"Last week no of orders":                   report.WeeklyOrders,
		"Last week totalAmount from orders":        report.WeeklyTotalAmount,
		"Last month no of orders":                  report.MonthlyOrders,
		"Last month totalAmount from orders":       report.MonthlyTotalAmount,
		"Last year no of orders":                   report.YearlyOrders,
		"Last year totalAmount from orders":        report.YearlyTotalAmount,
	})

}

// Delivery Status
//
//	@Summary		Updating Delivery Order By Admin
//	@Description	Admin Can Updates Delivery Status for Orders
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			orderItemID	path		string	true	"OrderItemID"
//	@Success		200			{string}	message
//	@Router			/v1/admin/updatingDeliveryStatus/{orderItemID} [put]
func (ac *AdminController) UpdatingDeliveryOrderByAdmin(c *gin.Context) {

	orderID, err := strconv.Atoi(c.Param("orderItemID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid orderID"})
		return
	}
	_, err1 := ac.adminService.UpdatingOrderStatus(uint(orderID))
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated sucessfully"})

}

// Return Status
//
//	@Summary		Updating Return Status ByAdmin
//	@Description	Admin Can Updates Return Status for Orders
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			orderItemID	path		string	true	"OrderItemID"
//	@Success		200			{string}	message
//	@Router			/v1/admin/updatingReturnStatus/{orderItemID} [put]
func (ac *AdminController) UpdatingReturnStatusByAdmin(c *gin.Context) {

	orderID, err := strconv.Atoi(c.Param("orderItemID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid orderID"})
		return
	}
	err = ac.adminService.UpdatingReturnOrderStatus(uint(orderID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "can't update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated sucessfully"})

}

// Add Coupon
//
//	@Summary		Add Coupon
//	@Description	Add Coupon
//	@Tags			Admin
//	@Param			coupon	body		deliverymodels.Coupon	true	"Coupon Data"
//	@Success		200		{string}	string					"Success Message"
//	@Router			/v1/admin/addCoupon [post]
func (ac *AdminController) AddCoupon(c *gin.Context) {

	var input deliverymodels.Coupon

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error binding data": err.Error()})
		return
	}
	var coupon entity.Coupon
	err := copier.Copy(&coupon, &input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = ac.adminService.AddCoupon(coupon)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coupon Added successfully"})

}

// LogOut
//
//	@Summary		Logout
//	@Description	Admin Logout,
//	@Tags			Admin
//	@Success		200	{string}	string	"Success Message"
//	@Router			/v1/admin/adminLogOut [post]
func (ac *AdminController) AdminLogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "70off.online", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
