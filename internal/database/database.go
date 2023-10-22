package database

import (
	"context"
	"fmt"

	"github.com/thomiaditya/shop-api/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Connection *gorm.DB
}

// Load the configuration from the environment variables and return a new PostgresDatabase instance
func NewPostgresDatabaseFromConfig() *PostgresDatabase {
	return &PostgresDatabase{
		Host:     util.GetEnv("DB_HOST", "localhost"),
		Port:     util.GetEnv("DB_PORT", "5432"),
		Username: util.GetEnv("DB_USERNAME", "postgres"),
		Password: util.GetEnv("DB_PASSWORD", "postgres"),
		Database: util.GetEnv("DB_DATABASE", "postgres"),
	}
}

func (db *PostgresDatabase) Connect(ctx context.Context) error {
	dsn := "host=" + db.Host + " user=" + db.Username + " password=" + db.Password + " dbname=" + db.Database + " port=" + db.Port + " sslmode=disable TimeZone=Asia/Jakarta"

	fmt.Println("Connecting to database...")

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("Connected to database:\nDB_HOST=" + db.Host + "\nDB_PORT=" + db.Port + "\nDB_USERNAME=" + db.Username + "\nDB_PASSWORD=" + db.Password + "\nDB_DATABASE=" + db.Database)

	db.Connection = conn
	return nil
}
