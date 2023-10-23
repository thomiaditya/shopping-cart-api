package model

import (
	"fmt"

	"gorm.io/gorm"
)

type CartStatus string

const (
	CartStatusOpen   CartStatus = "open"
	CartStatusClosed CartStatus = "closed"
)

type Cart struct {
	gorm.Model
	TotalAmount uint       `gorm:"type:integer;null"`
	Status      CartStatus `gorm:"type:varchar(100);default:open"`
	CustomerId  uint
	Customer    *Customer  `gorm:"foreignKey:CustomerId;references:ID"`
	Products    []*Product `gorm:"many2many:cart_items;"`
}

func (Cart) TableName() string {
	return "carts"
}

func NewCart(totalAmount uint, status CartStatus) *Cart {
	return &Cart{
		TotalAmount: totalAmount,
		Status:      status,
	}
}

func CreateNewCart(cart *Cart) (*Cart, error) {
	err := cart.Save()
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (cart *Cart) Save() error {
	return db.Create(cart).Error
}

// AddProduct adds a product to the cart. This method is used to add a product to the cart.
func (cart *Cart) AddProduct(product *Product, quantity uint) error {
	// Check if the product is already in the cart
	cartItem := &CartItem{
		CartId:    cart.ID,
		ProductId: product.ID,
		Quantity:  quantity,
	}

	err := cartItem.Save()
	if err != nil {
		return err
	}

	fmt.Println(db.Model(&cart).Association("Products").Count())

	return nil
}

// Close closes the cart.
func (cart *Cart) Close() error {
	cart.Status = CartStatusClosed
	return cart.Save()
}
