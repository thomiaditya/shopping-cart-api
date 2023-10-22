package database

import (
	"sync"

	"gorm.io/gorm"
)

type DatabaseInterface interface {
	Connect() error
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
