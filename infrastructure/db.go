package infrastructure

import (
	"70_Off/config"
	"70_Off/domain/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	gormDB *gorm.DB
}

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)
	fmt.Println("\ndsnnnn\n", dsn)

	if cfg.DB.Host == "" || cfg.DB.User == "" || cfg.DB.Password == "" || cfg.DB.Name == "" || cfg.DB.Port == "" {
		return nil, fmt.Errorf("missing database configuration values")
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := gormDB.AutoMigrate(
		&entity.User{},
		&entity.Address{},
		&entity.Admin{},
		&entity.Product{},
		&entity.ProductDetails{},
		&entity.Cart{},
		&entity.CartItem{},
		&entity.Orders{},
		&entity.OrderItems{},
		&entity.Invoice{},
		&entity.OtpValidation{},
		&entity.Brand{},
		&entity.Size{},
		&entity.Inventory{},
		&entity.Discount{},
		&entity.Payment{},
		&entity.PaymentType{},
		&entity.Wallet{},
		&entity.Coupon{},
	); err != nil {
		return nil, err
	}

	return gormDB, nil
}

// func (d *Database) User() repository.UserRepository {
// 	return repositoryImpl.NewUserRepositoryImpl(d.gormDB)
// }

// func (d *Database) Admin() repository.AdminRepository {
// 	return repositoryImpl.NewAdminRepositoryImpl(d.gormDB)
// }
