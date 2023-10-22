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
	ActiveCartId uint   `gorm:"type:integer;null"`
}

func (Customer) TableName() string {
	return "customers"
}

func NewCustomer(name string, email string, username string, passwordHash string) *Customer {
	return &Customer{
		Name:         name,
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
	}
}
