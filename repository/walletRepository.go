package repository

type WalletRepository interface {
	GetWalletByUserID(userID uint) (float64, error)
	GetWalletBalance(userID uint) (float64, error)
	DepositToWallet(userID uint, amount float64) error
	WithdrawFromWallet(userID uint, amount float64) error
}
