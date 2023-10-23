package main

import (
	"log"

	"github.com/thomiaditya/shop-api/internal/model"
)

func Seed() error {

	customer, err := model.CreateNewCustomer(&model.Customer{
		Name:         "Thomi Aditya",
		Email:        "thomiaditya@gmail.com",
		Username:     "thomiaditya",
		PasswordHash: "12345678",
	})
	if err != nil {
		return err
	}

	customer.AddCart(&model.Cart{
		TotalAmount: 100000000,
	})

	product, err := model.CreateNewProduct(&model.Product{
		Name:        "Macbook Pro 2020",
		Price:       20000000,
		Stock:       10,
		Description: "Macbook Pro 2020",
		Category:    "Laptop",
	})
	if err != nil {
		return err
	}

	customer.ActiveCart().AddProduct(product, 10)
	customer.ActiveCart().AddProduct(product, 10)

	return nil
}

func main() {
	log.Println("Seeding the database...")
	Seed()
	log.Println("Database seeded!")
}
