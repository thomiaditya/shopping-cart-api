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

func (product *Product) Update() error {
	return db.Save(product).Error
}

// Get the product based on the product ID.
func GetProduct(id uint) (*Product, error) {
	var product Product
	err := db.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Count the total number of products.
func CountProducts() (int64, error) {
	var count int64
	err := db.Model(&Product{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
