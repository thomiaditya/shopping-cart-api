package main

import (
	"log"

	"github.com/jaswdr/faker"
	"github.com/thomiaditya/shop-api/internal/model"
)

var fake = faker.New()

func SeedCustomers() error {

	customer, err := model.CreateNewCustomer(&model.Customer{
		Name:         fake.Person().Name(),
		Email:        fake.Internet().Email(),
		Username:     fake.Internet().User(),
		PasswordHash: "password",
	})
	if err != nil {
		return err
	}

	customer.CreateNewCart()

	return nil
}

func SeedProducts() error {
	_, err := model.CreateNewProduct(&model.Product{
		Name:        fake.Lorem().Word(),
		Price:       fake.UIntBetween(100, 1000),
		Stock:       10,
		Description: fake.Lorem().Sentence(10),
		Category:    fake.Lorem().Word(),
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	log.Println("Seeding the database...")
	for i := 0; i < 100; i++ {
		SeedCustomers()
		SeedProducts()
	}

	for i := 0; i < 10; i++ {
		customer, _ := model.GetCustomer(uint(fake.UIntBetween(1, 100)))
		product, _ := model.GetProduct(uint(fake.UIntBetween(1, 100)))

		cart := customer.ActiveCart()
		cart.AddProduct(product, uint(fake.UIntBetween(1, 10)))
	}

	log.Println("Database seeded!")
}
