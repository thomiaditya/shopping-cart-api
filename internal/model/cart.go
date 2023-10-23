package model

import (
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

	Customer  *Customer  `gorm:"foreignKey:CustomerId;references:ID"`
	Products  []*Product `gorm:"many2many:cart_items;"`
	OrderList []Order
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

// AddProduct adds a product to the cart. This method is used to add a product to the cart. If the product is already in the cart, the quantity of the product will be increased.
func (cart *Cart) AddProduct(product *Product, quantity uint) error {
	// Check if the product is already in the cart
	var cartItem CartItem
	err := db.Model(&CartItem{}).Where("cart_id = ? AND product_id = ?", cart.ID, product.ID).Limit(1).Find(&cartItem).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}

	// Update the cart item if it already exists
	if cartItem.ID != 0 {
		cartItem.Quantity += quantity
		err = cartItem.Update()
		if err != nil {
			return err
		}

		return nil
	}

	cartItem = CartItem{
		CartId:    cart.ID,
		ProductId: product.ID,
		Quantity:  quantity,
	}

	err = cartItem.Save()
	if err != nil {
		return err
	}

	return nil
}

// Close closes the cart.
func (cart *Cart) Close() error {
	cart.Status = CartStatusClosed
	return cart.Save()
}
