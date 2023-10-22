package database

import (
	"context"

	"gorm.io/gorm"
)

type DatabaseInterface interface {
	Connect(ctx context.Context) error
	GetConnection() *gorm.DB
}
