package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(100);not null"`
	Username     string `gorm:"type:varchar(100);not null"`
	PasswordHash string `gorm:"type:varchar(100);not null"`
	ActiveCartId uint

	CartList  []Cart
	OrderList []Order
}

func (Customer) TableName() string {
	return "customers"
}

func NewCustomer(name string, email string, username string, passwordHash string) *Customer {
	return &Customer{
		Name:         name,
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
	}
}

// CreateNewCustomer creates a new customer and saves it to the database.
func CreateNewCustomer(customer *Customer) (*Customer, error) {
	err := customer.Save()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// Get the customer based on the customer ID.
func GetCustomer(id uint) (*Customer, error) {
	var customer Customer
	err := db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// Save saves the customer to the database.
func (customer *Customer) Save() error {
	return db.Create(customer).Error
}

// Update updates the customer in the database.
func (customer *Customer) Update() error {
	return db.Save(customer).Error
}

// ActiveCart returns the active cart of the customer.
func (customer *Customer) ActiveCart() *Cart {
	var cart Cart
	db.Model(&Cart{}).Where("id = ?", customer.ActiveCartId).First(&cart)

	return &cart
}

// Carts returns all carts that belong to the customer.
func (customer *Customer) Carts() []Cart {
	var carts []Cart
	db.Model(&customer).Association("CartList").Find(&carts)
	return carts
}

// CreateNewCart adds a cart to the customer. This method is used to create a new cart for the customer and set it as the active cart.
func (customer *Customer) CreateNewCart() error {
	// Close previous active cart
	if customer.ActiveCartId != 0 {
		err := customer.ActiveCart().Close()
		if err != nil {
			return err
		}
	}

	var cart Cart
	err := db.Model(customer).Association("CartList").Append(&cart)
	if err != nil {
		return err
	}

	customer.ActiveCartId = cart.ID
	err = customer.Update()
	if err != nil {
		return err
	}

	return nil
}
