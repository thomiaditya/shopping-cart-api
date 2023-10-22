package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(100);not null"`
	Username     string `gorm:"type:varchar(100);not null"`
	PasswordHash string `gorm:"type:varchar(100);not null"`
	ActiveCartId uint   `gorm:"type:integer;not null"`
}

func (Customer) TableName() string {
	return "customers"
}
