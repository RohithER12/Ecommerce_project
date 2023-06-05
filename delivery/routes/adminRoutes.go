package routes

import (
	"70_Off/config"
	"70_Off/delivery/controllers"
	"70_Off/domain/auth"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine, adminController *controllers.AdminController, productController *controllers.ProductController) *gin.Engine {
	v1 := router.Group("/v1/admin")
	secret := config.GetJWTSecret()
	v1.Use(auth.AuthMiddleware(secret))

	router.POST("/adminSignup", adminController.AdminSignUp)
	router.POST("/adminLogin", adminController.AdminLogin)

	v1.POST("/addProduct", productController.AddProduct)
	v1.POST("/addProductDetails", productController.AddProductDetails)
	v1.PUT("/updateProduct", productController.UpdateProduct)
	v1.PUT("/updateProductDetails", productController.UpdateProductsDeatils)
	v1.PUT("/delete/:productId", productController.ProductSoftDelete)
	v1.GET("/listAllproducts", productController.ListProducts)
	v1.PUT("/userBlockToggle/:userID", adminController.UserBLockToggle)
	v1.POST("/adminLogOut", adminController.AdminLogoutHandler)
	v1.GET("/saleReport/:days", adminController.OrderReport)
	v1.PUT("/updatingDeliveryStatus/:orderItemID", adminController.UpdatingDeliveryOrderByAdmin)
	v1.PUT("/updatingReturnStatus/:orderItemID", adminController.UpdatingReturnStatusByAdmin)
	v1.POST("/addCoupon", adminController.AddCoupon)

	return router
}
