package migration

import (
	"context"
	"log"

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

func (m *Migrate) ConnectDatabase(ctx context.Context) (*Migrate, error) {
	// Load the configuration
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Connect to the database
	m.db = database.GetDatabaseInstance()

	return m, nil
}

func (m *Migrate) Up(ctx context.Context) error {
	log.Println("Migrating database...")

	// Migrate the database
	modelList := m.modelsToMigrate(m.db.GetConnection())
	for _, model := range modelList {
		err := m.db.GetConnection().AutoMigrate(model)
		if err != nil {
			return err
		}
		log.Printf("Migrated model: %T\n", model)
	}

	log.Println("Migrated database")

	return nil
}

func (m *Migrate) Down(ctx context.Context) error {
	log.Println("Rolling back database...")

	// Rollback the database
	modelList := m.modelsToMigrate(m.db.GetConnection())
	for _, model := range modelList {
		err := m.db.GetConnection().Migrator().DropTable(model)
		if err != nil {
			return err
		}
		log.Printf("Rolled back model: %T\n", model)
	}

	log.Println("Rolled back database")

	return nil
}
