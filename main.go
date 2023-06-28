package main

import (
	"70_Off/delivery/controllers"
	"70_Off/delivery/routes"
	_ "70_Off/docs"
	"70_Off/infrastructure"
	"70_Off/infrastructure/config"
	"70_Off/repository/repositoryImpl"
	"70_Off/services"
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title						70_Off
// @version					1.0
// @description				Testing Swagger APIs.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						token
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8088
// @BasePath					/api/v1
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	db, err := infrastructure.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Error creating database connection:", err)
	}

	userRepo := repositoryImpl.NewUserRepositoryImpl(db)
	adminRepo := repositoryImpl.NewAdminRepositoryImpl(db)
	productRepo := repositoryImpl.NewProductRepositoryImpl(db)
	cartRepo := repositoryImpl.NewCartRepoImpl(db)
	cartItemRepo := repositoryImpl.NewCartItemRepoImpl(db)
	orderRepo := repositoryImpl.NewOrderRepoImpl(db)
	paymentRepo := repositoryImpl.NewPaymentRepo(db)
	invoiceRepo := repositoryImpl.NewInvoiceRepo(db)
	categoryRepo := repositoryImpl.NewCategoryRepositoryImpl(db)
	otpValidationRepo := repositoryImpl.NewOtpValidationRepoImpl(db)
	inventoryRepo := repositoryImpl.NewInventoryRepositoryImpl(db)
	walletRepo := repositoryImpl.NewWalletRepository(db)
	couponRepo := repositoryImpl.NewCouponRepoImpl(db)

	userService := services.NewUserService(userRepo, cartRepo, cartItemRepo, orderRepo,
		invoiceRepo, otpValidationRepo, inventoryRepo, productRepo, paymentRepo, walletRepo, couponRepo)
	adminSerivice := services.NewAdminService(adminRepo, userRepo, inventoryRepo, orderRepo, walletRepo, couponRepo)
	productSerivice := services.NewProductService(productRepo, categoryRepo, inventoryRepo, orderRepo)

	userController := controllers.NewUserController(userService)
	adminController := controllers.NewAdminController(adminSerivice)
	productController := controllers.NewProductController(*productSerivice)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.UserRouter(router, userController, productController)
	routes.AdminRouter(router, adminController, productController)

	router.Run(":" + cfg.Server.Address)
}
