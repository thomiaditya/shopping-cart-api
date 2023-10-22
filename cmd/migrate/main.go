package main

import (
	"context"
	"fmt"
	"os"

	"github.com/thomiaditya/shop-api/internal/migration"
)

func main() {
	var argument string
	if len(os.Args) > 1 {
		argument = os.Args[1]
	}

	// Create a new migrate instance
	migrate := migration.NewMigrate()

	// Start the migrate process
	migration, err := migrate.ConnectDatabase(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if argument == "down" {
		_ = migration.Down(context.Background())
		return
	}

	migration.Up(context.Background())
}
