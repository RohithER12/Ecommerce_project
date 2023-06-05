package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	HouseNo    string `json:"houseNo"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	UserID     uint   `json:"userId"`
	User       User   `gorm:"foreignkey:UserID" json:"user"`
}
