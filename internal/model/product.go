package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null"`
	Price       uint    `gorm:"type:integer;not null"`
	Stock       uint    `gorm:"type:integer;not null"`
	Description string  `gorm:"type:varchar(100);not null"`
	Category    string  `gorm:"type:varchar(100);not null"`
	Cart        []*Cart `gorm:"many2many:cart_items;"`
}

func (Product) TableName() string {
	return "products"
}

func NewProduct(name string, price uint, stock uint, description string, category string) *Product {
	return &Product{
		Name:        name,
		Price:       price,
		Stock:       stock,
		Description: description,
		Category:    category,
	}
}

func CreateNewProduct(product *Product) (*Product, error) {
	err := product.Save()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (product *Product) Save() error {
	return db.Create(product).Error
}
