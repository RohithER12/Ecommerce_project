package entity

import (
	"gorm.io/gorm"
)

type Size struct {
	gorm.Model
	Size string `json:"size"`
}

type Brand struct {
	gorm.Model
	Brand string `json:"brand"`
}
