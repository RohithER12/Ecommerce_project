package routes

import (
	"70_Off/delivery/auth"
	"70_Off/delivery/controllers"
	"70_Off/infrastructure/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, userController *controllers.UserController, productController *controllers.ProductController) *gin.Engine {
	// Initialize session middleware
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	v1 := router.Group("/v1")
	secret := config.GetJWTSecret()

	v1.Use(auth.AuthMiddleware(secret))

	router.POST("/userSignup", userController.UserSignUp)
	router.POST("/userSignupVerification", userController.UserSignUpVerification)
	router.POST("/userLogin", userController.UserLogin)
	router.POST("/userLoginWithOtp", userController.UserLoginWithOtp)
	router.POST("/userLoginWithOtpValidation", userController.UserLoginWithOtpValidation)
	router.GET("/resetPasswordRequst", userController.ResetPasswordRequestUser)
	router.PUT("/resetPassword", userController.ResetPasswordUser)

	v1.POST("/addAddress", userController.AddAddress)
	v1.PUT("/updateUserDetails", userController.UserUpdates)
	v1.POST("/addCart/:product/:quantity/:sizeID", userController.AddCartItem)
	v1.GET("/diplayCart", userController.DisplayCart)
	v1.PUT("/removeCartItem/:productItemID/:quantity", userController.RemoveCartItem)
	v1.GET("/orderProduct", userController.CheckOutCart)
	v1.POST("/CheckOutCartProcess/:paymentTypeID/:addressID/:coupon", userController.PaymentSelection)
	v1.POST("razorpayVerification/:sign/:razorid/:payid", userController.RazorPayVerification)
	v1.GET("/orderHistory", userController.OrderHistory)
	v1.PUT("/cancelOrderByUser/:orderID", userController.CancelOrderByUser)
	v1.PUT("/returnOrderByUser/:orderID", userController.ReturnOrderByUser)
	v1.GET("/productsPagination", productController.ListProductsPagination)
	v1.GET("/filterByBrand/:brandID", userController.ProductCategoryFiltering)
	v1.GET("/searchByName/:search", userController.SearchProductsByName)
	v1.GET("/walletBalance", userController.DisplayWalletBalance)
	v1.GET("/productDetail/:productID", userController.ProductDetails)
	v1.GET("/userProfile", userController.UserProfile)
	v1.POST("/logOut", userController.LogoutHandler)

	return router
}
