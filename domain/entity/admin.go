package entity

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber" gorm:"unique_index"`
	Password    string `json:"password"`
}
