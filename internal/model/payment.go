package model

import "gorm.io/gorm"

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
)

type Payment struct {
	gorm.Model
	Method string        `gorm:"type:varchar(100);not null"`
	Status PaymentStatus `gorm:"type:varchar(100);default:pending"`
	Amount uint          `gorm:"type:integer;not null"`
}
