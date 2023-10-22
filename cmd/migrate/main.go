package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/thomiaditya/shop-api/internal/database"
	"github.com/thomiaditya/shop-api/internal/model"
	"gorm.io/gorm"
)

type Migrate struct {
	db database.DatabaseInterface
}

func NewMigrate() *Migrate {
	return &Migrate{}
}

func (m *Migrate) modelsToMigrate(db *gorm.DB) []interface{} {
	return []interface{}{
		&model.Customer{},
	}
}

func (m *Migrate) Start(ctx context.Context) error {
	// Load the configuration
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Connect to the database
	m.db = database.NewPostgresDatabaseFromConfig()
	err = m.db.Connect(ctx)
	if err != nil {
		return err
	}

	// Migrate the database
	err = m.migrate(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *Migrate) migrate(ctx context.Context) error {
	fmt.Println("Migrating database...")

	// Migrate the database
	modelList := m.modelsToMigrate(m.db.GetConnection())
	for _, model := range modelList {
		err := m.db.GetConnection().AutoMigrate(model)
		if err != nil {
			return err
		}
		fmt.Printf("Migrated model: %T\n", model)
	}

	fmt.Println("Migrated database")

	return nil
}

func main() {
	// Create a new migrate instance
	migrate := NewMigrate()

	// Start the migrate process
	err := migrate.Start(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
