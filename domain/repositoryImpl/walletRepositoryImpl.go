package repositoryImpl

import (
	"70_Off/domain/entity"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type walletRepositoryImpl struct {
	db *gorm.DB
	mu sync.RWMutex
}

// func NewWalletRepository(db *gorm.DB) *walletRepositoryImpl {
// 	return &walletRepositoryImpl{db}

// }
func NewWalletRepository(db *gorm.DB) *walletRepositoryImpl {
	return &walletRepositoryImpl{
		db: db,
		mu: sync.RWMutex{},
	}
}

func (wr *walletRepositoryImpl) GetWalletByUserID(userID uint) (float64, error) {
	wr.mu.RLock()
	defer wr.mu.RUnlock()

	var wallet entity.Wallet
	if err := wr.db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("wallet not found for user ID %d", userID)
		}
		return 0, err
	}

	return wallet.Balance, nil
}

func (wr *walletRepositoryImpl) GetWalletBalance(userID uint) (float64, error) {
	return wr.GetWalletByUserID(userID)
}

func (wr *walletRepositoryImpl) DepositToWallet(userID uint, amount float64) error {
	wr.mu.Lock()
	defer wr.mu.Unlock()

	if err := wr.db.Transaction(func(tx *gorm.DB) error {
		var wallet entity.Wallet
		if err := tx.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create a new wallet if it doesn't exist for the user
				wallet = entity.Wallet{UserID: userID, Balance: amount}
				return tx.Create(&wallet).Error
			}
			return err
		}

		// Update the wallet balance by adding the deposited amount
		wallet.Balance += amount
		return tx.Save(&wallet).Error
	}); err != nil {
		return err
	}

	return nil
}

func (wr *walletRepositoryImpl) WithdrawFromWallet(userID uint, amount float64) error {
	wr.mu.Lock()
	defer wr.mu.Unlock()

	if userID == 0 {
		return fmt.Errorf("invalid user ID")
	}

	if amount <= 0 {
		return fmt.Errorf("invalid withdrawal amount")
	}

	var wallet entity.Wallet
	if err := wr.db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("wallet not found for user ID %d", userID)
		}
		return fmt.Errorf("failed to fetch wallet for user ID %d: %w", userID, err)
	}

	if wallet.Balance < amount {
		return fmt.Errorf("insufficient funds in the wallet")
	}

	wallet.Balance -= amount
	if err := wr.db.Model(&wallet).Where("user_id = ?", userID).Update("balance", wallet.Balance).Error; err != nil {
		return fmt.Errorf("failed to update wallet: %w", err)
	}

	return nil
}
