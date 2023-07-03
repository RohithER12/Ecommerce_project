package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `json:"firstname"  `
	LastName    string `json:"lastname" `
	Email       string `json:"email"  gorm:"unique_index"`
	PhoneNumber string `json:"phonenumber" gorm:"unique_index"`
	Username    string `json:"username" gorm:"unique_index"`
	Password    string `json:"password" `
	Block       bool   `gorm:"not null;default:false"`
	Verified    bool   `gorm:"not null;default:false"`
}

type UserLoginWithOtp struct {
	PhoneNumber string `json:"phonenumber"`
	Key         string `json:"key"`
	Otp         string `json:"otp"`
	ResendOtp   bool   `gorm:"not null;default:false"`
	Password    string `json:"password"`
}
