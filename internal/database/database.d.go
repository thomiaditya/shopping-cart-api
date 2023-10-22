package database

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

type DatabaseInterface interface {
	Connect(ctx context.Context) error
	GetConnection() *gorm.DB
}

var instance DatabaseInterface
var once sync.Once

// Return a singleton instance of PostgresDatabase
func GetDatabaseInstance() DatabaseInterface {
	once.Do(func() {
		instance = NewPostgresDatabaseFromConfig()
	})
	return instance
}
