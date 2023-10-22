package main

import (
	"log"

	"github.com/thomiaditya/shop-api/internal/model"
)

func Seed() error {
	// Create the customers
	customers := []model.Customer{
		{
			Name:         "John Doe",
			Email:        "johndoe@gmail.com",
			Username:     "johndoe",
			PasswordHash: "123456",
		},
	}

	for _, customer := range customers {
		err := customer.Save()
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	log.Println("Seeding the database...")
	Seed()
	log.Println("Database seeded!")
}
