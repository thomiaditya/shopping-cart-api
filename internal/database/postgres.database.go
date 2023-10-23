package database

import (
	"log"

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
	db := &PostgresDatabase{
		Host:     util.GetEnv("DB_HOST", "localhost"),
		Port:     util.GetEnv("DB_PORT", "5432"),
		Username: util.GetEnv("DB_USERNAME", "postgres"),
		Password: util.GetEnv("DB_PASSWORD", "postgres"),
		Database: util.GetEnv("DB_DATABASE", "postgres"),
	}

	db.Connect()
	return db
}

func (db *PostgresDatabase) Connect() error {
	dsn := "host=" + db.Host + " user=" + db.Username + " password=" + db.Password + " dbname=" + db.Database + " port=" + db.Port + " sslmode=disable TimeZone=Asia/Jakarta"

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Database connected!")

	db.Connection = conn
	return nil
}

func (db *PostgresDatabase) GetConnection() *gorm.DB {
	return db.Connection
}
